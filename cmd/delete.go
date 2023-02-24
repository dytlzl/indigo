package cmd

import (
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
	RunE: func(cmd *cobra.Command, args []string) error {
		return instanceUsecase.Delete(cmd.Context(), args[0])
	},
}

var deleteFirewallCmd = &cobra.Command{
	Use:     "firewall [name]",
	Aliases: []string{"fw", "firewalls"},
	Short:   "Delete a firewall",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return firewallUsecase.Delete(cmd.Context(), args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteInstanceCmd)
	deleteCmd.AddCommand(deleteFirewallCmd)
	rootCmd.AddCommand(deleteCmd)
}
