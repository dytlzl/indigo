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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop [name]",
	Short: "Stop an instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client, err := api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
		}
		repo := repository.NewAPIInstanceRepository(client)
		u := usecase.NewInstanceUsecase(repo)
		force, err := cmd.PersistentFlags().GetBool("force")
		if err != nil {
			log.Fatalln(err)
		}
		if force {
			err = u.ForceStop(ctx, args[0])
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			err = u.Stop(ctx, args[0])
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	stopCmd.PersistentFlags().BoolP("force", "f", false, "indigo stop node00 --force")
	rootCmd.AddCommand(stopCmd)
}
