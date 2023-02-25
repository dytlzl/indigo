package cmd

import (
	"context"
	"log"
	"os"
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
func Execute() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := NewRootCmd().ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

var (
	client          api.Client
	instanceUseCase usecase.InstanceUseCase
	firewallUseCase usecase.FirewallUseCase
	osUseCase       usecase.OSUseCase
	sshKeyUseCase   usecase.SSHKeyUseCase
	planUseCase     usecase.PlanUseCase
)

func NewRootCmd() *cobra.Command {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	cmd := &cobra.Command{
		Use:   "indigo",
		Short: "Indigo API client",
		Long:  "indigo is a Indigo API client written in Go.",
	}
	var configFile string
	cmd.PersistentFlags().StringVar(&configFile, "config", filepath.Join(u.HomeDir, ".indigo.yaml"), "config file (default is $HOME/.indigo.yaml)")
	cobra.OnInitialize(func() {
		conf := config.NewConfig(configFile)
		client, err = api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
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
	})
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
	return cmd
}
