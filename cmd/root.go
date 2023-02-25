package cmd

import (
	"context"
	"os/user"
	"path/filepath"
	"time"

	"github.com/dytlzl/indigo/pkg/config"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/infra/repository"
	"github.com/dytlzl/indigo/pkg/usecase"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := &cobra.Command{
		Use:   "indigo",
		Short: "Indigo API client",
		Long:  "indigo is a Indigo API client written in Go.",
	}
	var configFilename string
	u, err := user.Current()
	if err != nil {
		return err
	}
	cmd.PersistentFlags().StringVar(&configFilename, "config", filepath.Join(u.HomeDir, ".indigo.yaml"), "config file (default is $HOME/.indigo.yaml)")
	cmd.AddCommand(
		NewApplyCmd(),
		NewCreatCmd(),
		NewDeleteCmd(),
		NewGetCmd(),
		NewStartCmd(),
		NewStopCmd(),
		versionCmd,
		debugCmd,
	)
	err = InitUseCases(configFilename)
	if err != nil {
		return err
	}
	err = cmd.ExecuteContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

var (
	client          api.Client
	instanceUseCase usecase.InstanceUseCase
	firewallUseCase usecase.FirewallUseCase
	osUseCase       usecase.OSUseCase
	sshKeyUseCase   usecase.SSHKeyUseCase
	planUseCase     usecase.PlanUseCase
)

func InitUseCases(configFilename string) error {
	conf := config.NewConfig(configFilename)
	client, err := api.NewClient(conf)
	if err != nil {
		return err
	}
	ir := repository.NewAPIInstanceRepository(client)
	fr := repository.NewAPIFirewallRepository(client)
	or := repository.NewAPIOSRepository(client)
	pr := repository.NewJSONPlanRepository()
	sr := repository.NewAPISSHKeyRepository(client)
	instanceUseCase = usecase.NewInstanceUseCase(ir)
	firewallUseCase = usecase.NewFirewallUseCase(fr, ir)
	osUseCase = usecase.NewOSUseCase(or)
	sshKeyUseCase = usecase.NewSSHKeyUseCase(sr)
	planUseCase = usecase.NewPlanUseCase(pr)
	return nil
}
