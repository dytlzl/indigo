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
	"github.com/dytlzl/indigo/pkg/repository"
	"github.com/dytlzl/indigo/pkg/usecase"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "indigo",
	Short: "Indigo API client",
	Long:  "indigo is a Indigo API client written in Go.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

var configFile string

var conf config.Config

var instanceUsecase usecase.InstanceUsecase

var firewallUsecase usecase.FirewallUsecase

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	rootCmd.PersistentFlags().StringVar(&configFile, "config", filepath.Join(u.HomeDir, ".indigo.yaml"), "config file (default is $HOME/.indigo.yaml)")
	cobra.OnInitialize(func() {
		conf = config.NewConfig(configFile)
		client, err := api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
		}
		ir := repository.NewAPIInstanceRepository(client)
		fr := repository.NewAPIFirewallRepository(client)
		instanceUsecase = usecase.NewInstanceUsecase(ir)
		firewallUsecase = usecase.NewFirewallUsecase(fr, ir)
	})
}
