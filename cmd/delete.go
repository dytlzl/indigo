package cmd

import (
	"github.com/dytlzl/indigo/cmd/di"
	"github.com/dytlzl/indigo/pkg/config"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(conf config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a resource",
	}
	deleteInstanceCmd := &cobra.Command{
		Use:     "instance [name]",
		Aliases: []string{"i", "instances"},
		Short:   "Delete an instance",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return di.InitializeInstanceUseCase(conf).Delete(cmd.Context(), args[0])
		},
	}
	deleteFirewallCmd := &cobra.Command{
		Use:     "firewall [name]",
		Aliases: []string{"fw", "firewalls"},
		Short:   "Delete a firewall",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return di.InitializeFirewallUseCase(conf).Delete(cmd.Context(), args[0])
		},
	}
	cmd.AddCommand(
		deleteInstanceCmd,
		deleteFirewallCmd,
	)
	return cmd
}
