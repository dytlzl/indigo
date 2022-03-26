package cmd

import (
	"context"
	"log"
	"time"

	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/repository"
	"github.com/dytlzl/indigo/pkg/usecase"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource(s)",
}

var getInstanceCmd = &cobra.Command{
	Use:     "instance",
	Aliases: []string{"i", "instances"},
	Short:   "Get instance(s)",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client, err := api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
		}
		repo := repository.NewAPIInstanceRepository(client)
		u := usecase.NewInstanceUsecase(repo)
		err = u.List(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

// getFirewallCmd represents the firewall command
var getFirewallCmd = &cobra.Command{
	Use:     "firewall",
	Aliases: []string{"fw", "firewalls"},
	Short:   "Get firewall(s)",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client, err := api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
		}
		repo := repository.NewAPIFirewallRepository(client)
		u := usecase.NewFirewallUsecase(repo, nil)
		if len(args) != 0 {
			err = u.Get(ctx, args[0])
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			err = u.List(ctx)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	getCmd.AddCommand(getInstanceCmd)
	getCmd.AddCommand(getFirewallCmd)
	rootCmd.AddCommand(getCmd)
}
