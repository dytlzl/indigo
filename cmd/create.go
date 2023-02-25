package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func NewCreatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a resource",
	}
	var (
		planID   int
		osID     int
		regionID int
		sshKeyID int
	)
	var createInstanceCmd = &cobra.Command{
		Use:     "instance [name]",
		Aliases: []string{"i", "instances"},
		Short:   "Create a instance",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return instanceUseCase.Create(cmd.Context(), args[0], planID, osID, regionID, sshKeyID)
		},
	}
	createInstanceCmd.Flags().IntVar(&planID, "plan-id", -1, "plan(Size) ID")
	err := createInstanceCmd.MarkFlagRequired("plan-id")
	if err != nil {
		log.Fatalln(err)
	}
	createInstanceCmd.Flags().IntVar(&sshKeyID, "ssh-key-id", -1, "SSH key ID")
	err = createInstanceCmd.MarkFlagRequired("ssh-key-id")
	if err != nil {
		log.Fatalln(err)
	}
	createInstanceCmd.Flags().IntVar(&osID, "os-id", 13, "OS ID (Ubuntu 22.04 13)")
	createInstanceCmd.Flags().IntVar(&regionID, "region-id", 1, "region ID")
	cmd.AddCommand(createInstanceCmd)
	return cmd
}
