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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource",
}

var deleteInstanceCmd = &cobra.Command{
	Use:     "instance [name]",
	Aliases: []string{"i", "instances"},
	Short:   "Delete an instance",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client, err := api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
		}
		repo := repository.NewAPIInstanceRepository(client)
		u := usecase.NewInstanceUsecase(repo)
		err = u.Delete(ctx, args[0])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var deleteFirewallCmd = &cobra.Command{
	Use:     "firewall [name]",
	Aliases: []string{"fw", "firewalls"},
	Short:   "Delete a firewall",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client, err := api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
		}
		repo := repository.NewAPIFirewallRepository(client)
		u := usecase.NewFirewallUsecase(repo, nil)
		err = u.Delete(ctx, args[0])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteInstanceCmd)
	deleteCmd.AddCommand(deleteFirewallCmd)
	rootCmd.AddCommand(deleteCmd)
}
