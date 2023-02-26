package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/usecase"
)

type apiFirewallRepository struct {
	Client api.Client
}

type firewallListResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ServiceID string `json:"service_id"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserID    int    `json:"user_id"`
}

func NewAPIFirewallRepository(client api.Client) usecase.FirewallRepository {
	return &apiFirewallRepository{Client: client}
}

func (a *apiFirewallRepository) List(ctx context.Context) ([]domain.Firewall, error) {
	bytes, err := a.Client.Get(ctx, "/nw/getfirewalllist")
	if err != nil {
		return nil, err
	}
	firewallList := make([]firewallListResponse, 0)
	err = json.Unmarshal(bytes, &firewallList)
	if err != nil {
		return nil, err
	}
	firewalls := make([]domain.Firewall, len(firewallList))
	for index, firewall := range firewallList {
		firewalls[index].ID = firewall.ID
		firewalls[index].Name = firewall.Name
		t, _ := time.Parse("2006-01-02 15:04:05", firewall.CreatedAt)
		firewalls[index].CreatedAt = t
	}
	return firewalls, nil
}

type templateResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Direction string `json:"direction"`
	Port      string `json:"port"`
	Protocol  string `json:"protocol"`
	Source    string `json:"source"`
}

func (a *apiFirewallRepository) Get(ctx context.Context, id int) (domain.Firewall, error) {
	bytes, err := a.Client.Get(ctx, fmt.Sprintf("/nw/gettemplate/%d", id))
	if err != nil {
		return domain.Firewall{}, err
	}
	templateList := make([]templateResponse, 0)
	err = json.Unmarshal(bytes, &templateList)
	if err != nil {
		return domain.Firewall{}, err
	}
	if len(templateList) == 0 {
		return domain.Firewall{}, nil
	}
	firewall := domain.Firewall{}
	firewall.ID = id
	firewall.Name = templateList[0].Name
	firewall.Inbound = make([]domain.Rule, 0, len(templateList))
	firewall.Outbound = make([]domain.Rule, 0, len(templateList))
	for _, template := range templateList {
		if template.Direction == "in" {
			firewall.Inbound = append(firewall.Inbound, domain.Rule{Type: template.Type, Protocol: template.Protocol, Port: template.Port, Source: template.Source})
		} else if template.Direction == "out" {
			firewall.Outbound = append(firewall.Outbound, domain.Rule{Type: template.Type, Protocol: template.Protocol, Port: template.Port, Source: template.Source})
		}
	}
	return firewall, nil
}

type FirewallRequest struct {
	TemplateID int           `json:"templateid"`
	Name       string        `json:"name"`
	Inbound    []domain.Rule `json:"inbound"`
	Outbound   []domain.Rule `json:"outbound"`
	Instances  []string      `json:"instances"`
}

func (a *apiFirewallRepository) Update(ctx context.Context, fw domain.Firewall) error {
	firewallRequest := FirewallRequest{}
	firewallRequest.TemplateID = fw.ID
	firewallRequest.Name = fw.Name
	firewallRequest.Inbound = fw.Inbound
	firewallRequest.Outbound = fw.Outbound
	firewallRequest.Instances = fw.Instances
	reqBody, err := json.Marshal(firewallRequest)
	if err != nil {
		return err
	}
	resBody, err := a.Client.Put(ctx, "/nw/updatefirewall", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	log.Println(string(resBody))
	return nil
}

func (a *apiFirewallRepository) Create(ctx context.Context, fw domain.Firewall) error {
	firewallRequest := FirewallRequest{}
	firewallRequest.TemplateID = fw.ID
	firewallRequest.Name = fw.Name
	firewallRequest.Inbound = fw.Inbound
	firewallRequest.Outbound = fw.Outbound
	firewallRequest.Instances = fw.Instances
	reqBody, err := json.Marshal(firewallRequest)
	if err != nil {
		return err
	}
	resBody, err := a.Client.Post(ctx, "/nw/createfirewall", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	log.Println(string(resBody))
	return nil
}

func (a *apiFirewallRepository) Delete(ctx context.Context, id int) error {
	bytes, err := a.Client.Delete(ctx, fmt.Sprintf("/nw/deletefirewall/%d", id))
	log.Println(string(bytes))
	return err
}
