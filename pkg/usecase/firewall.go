package usecase

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/printutil"

	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/util/duration"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_$GOPACKAGE

type FirewallRepository interface {
	List(ctx context.Context) ([]domain.Firewall, error)
	Get(ctx context.Context, id int) (*domain.Firewall, error)
	Create(ctx context.Context, fw *domain.Firewall) error
	Update(ctx context.Context, fw *domain.Firewall) error
	Delete(ctx context.Context, id int) error
}

type FirewallUseCase interface {
	List(ctx context.Context) error
	Get(ctx context.Context, target string) error
	Apply(ctx context.Context, fileBody []byte) error
	Delete(ctx context.Context, target string) error
}

type firewallUseCase struct {
	firewallRepository FirewallRepository
	instanceRepository InstanceRepository
}

func NewFirewallUseCase(fr FirewallRepository, ir InstanceRepository) FirewallUseCase {
	return &firewallUseCase{
		firewallRepository: fr,
		instanceRepository: ir,
	}
}

func (u *firewallUseCase) List(ctx context.Context) error {
	firewalls, err := u.firewallRepository.List(ctx)
	if err != nil {
		return err
	}
	printutil.PrintTable(
		[]string{"NAME", "AGE"},
		firewalls,
		func(firewall domain.Firewall) []string {
			return []string{firewall.Name, duration.HumanDuration(time.Since(firewall.CreatedAt))}
		},
		"",
	)
	return nil
}

func (u *firewallUseCase) Get(ctx context.Context, target string) error {
	fws, err := u.firewallRepository.List(ctx)
	if err != nil {
		return err
	}
	for _, element := range fws {
		if target == element.Name {
			firewall, err := u.firewallRepository.Get(ctx, element.ID)
			if err != nil {
				return err
			}
			fmt.Printf("ID: %d\n", element.ID)
			fmt.Printf("NAME: %s\n", element.Name)
			fmt.Println("INBOUND:")
			printutil.PrintTable(
				[]string{"TYPE", "PROTOCOL", "PORT", "SOURCE"},
				firewall.Inbound,
				func(rule domain.Rule) []string {
					return []string{rule.Type, rule.Protocol, rule.Port, rule.Source}
				},
				"  ",
			)
			fmt.Println("OUTBOUND:")
			printutil.PrintTable(
				[]string{"TYPE", "PROTOCOL", "PORT", "SOURCE"},
				firewall.Outbound,
				func(rule domain.Rule) []string {
					return []string{rule.Type, rule.Protocol, rule.Port, rule.Source}
				},
				"  ",
			)
			return nil
		}
	}
	return fmt.Errorf("firewall \"%s\" not found", target)
}

func (u *firewallUseCase) Apply(ctx context.Context, fileBody []byte) error {
	fw := domain.Firewall{}
	err := yaml.Unmarshal(fileBody, &fw)
	if err != nil {
		return err
	}
	fw.ID = 0
	fws, err := u.firewallRepository.List(ctx)
	if err != nil {
		return err
	}
	for _, element := range fws {
		if fw.Name == element.Name {
			fw.ID = element.ID
		}
	}
	instanceNameMap, err := u.getInstanceNameMap(ctx)
	if err != nil {
		return err
	}
	for index, name := range fw.Instances {
		instance, ok := instanceNameMap[name]
		if ok {
			fw.Instances[index] = strconv.Itoa(instance.ID)
		} else {
			return fmt.Errorf("invalid instance name \"%s\" was found in instances", name)
		}
	}
	for index, rule := range fw.Inbound {
		_, _, err := net.ParseCIDR(rule.Source)
		if err != nil {
			instance, ok := instanceNameMap[rule.Source]
			if ok {
				fw.Inbound[index].Source = instance.IP + "/32"
			} else {
				return fmt.Errorf("invalid source \"%s\" was found in inbounds", rule.Source)
			}
		}
	}
	for index, rule := range fw.Outbound {
		_, _, err := net.ParseCIDR(rule.Source)
		if err != nil {
			instance, ok := instanceNameMap[rule.Source]
			if ok {
				fw.Outbound[index].Source = instance.IP + "/32"
			} else {
				return fmt.Errorf("invalid source \"%s\" was found in inbounds", rule.Source)
			}
		}
	}
	if fw.ID == 0 {
		err = u.firewallRepository.Create(ctx, &fw)
		if err != nil {
			return err
		}
		fmt.Printf("firewall \"%s\" created\n", fw.Name)
	} else {
		err = u.firewallRepository.Update(ctx, &fw)
		if err != nil {
			return err
		}
		fmt.Printf("firewall \"%s\" configured\n", fw.Name)
	}
	return nil
}

func (u *firewallUseCase) getInstanceNameMap(ctx context.Context) (map[string]domain.Instance, error) {
	instances, err := u.instanceRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	instanceNameMap := map[string]domain.Instance{}
	for _, element := range instances {
		instanceNameMap[element.Name] = element
	}
	return instanceNameMap, nil
}

func (u *firewallUseCase) Delete(ctx context.Context, target string) error {
	fws, err := u.firewallRepository.List(ctx)
	if err != nil {
		return err
	}
	for _, element := range fws {
		if target == element.Name {
			err = u.firewallRepository.Delete(ctx, element.ID)
			if err != nil {
				return err
			}
			fmt.Printf("firewall \"%s\" deleted\n", target)
			return nil
		}
	}
	return fmt.Errorf("firewall \"%s\" not found", target)
}
