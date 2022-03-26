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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [name]",
	Short: "Start an instance",
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
		err = u.Start(ctx, args[0])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
