package cmd

import (
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
	RunE: func(cmd *cobra.Command, args []string) error {
		return instanceUseCase.Create(cmd.Context(), args[0], *planID, *osID, *regionID, *sshKeyID)
	},
}
var (
	planID   *int
	osID     *int
	regionID *int
	sshKeyID *int
)

func init() {
	planID = createInstanceCmd.Flags().Int("plan-id", -1, "plan(Size) ID")
	osID = createInstanceCmd.Flags().Int("os-id", 13, "OS ID (Ubuntu 22.04 13)")
	regionID = createInstanceCmd.Flags().Int("region-id", 1, "region ID")
	sshKeyID = createInstanceCmd.Flags().Int("ssh-key-id", -1, "SSH key ID")
	createInstanceCmd.MarkFlagRequired("plan-id")
	createInstanceCmd.MarkFlagRequired("ssh-key-id")
	createCmd.AddCommand(createInstanceCmd)
}
