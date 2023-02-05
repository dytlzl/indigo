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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
}

var createInstanceCmd = &cobra.Command{
	Use:     "instance [name]",
	Aliases: []string{"i", "instances"},
	Short:   "Create a instance",
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
		planID, err := cmd.Flags().GetInt("plan-id")
		if err != nil {
			log.Fatalln(err)
		}
		osID, err := cmd.Flags().GetInt("os-id")
		if err != nil {
			log.Fatalln(err)
		}
		regionID, err := cmd.Flags().GetInt("region-id")
		if err != nil {
			log.Fatalln(err)
		}
		sshKeyID, err := cmd.Flags().GetInt("ssh-key-id")
		if err != nil {
			log.Fatalln(err)
		}
		err = u.Create(ctx, args[0], planID, osID, regionID, sshKeyID)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	createInstanceCmd.Flags().Int("plan-id", -1, "plan(Size) ID")
	createInstanceCmd.Flags().Int("os-id", 13, "OS ID (Ubuntu 22.04 13)")
	createInstanceCmd.Flags().Int("region-id", 1, "region ID")
	createInstanceCmd.Flags().Int("ssh-key-id", -1, "SSH key ID")
	createInstanceCmd.MarkFlagRequired("plan-id")
	createInstanceCmd.MarkFlagRequired("ssh-key-id")
	createCmd.AddCommand(createInstanceCmd)
	rootCmd.AddCommand(createCmd)
}
