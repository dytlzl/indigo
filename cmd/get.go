package cmd

import (
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
	RunE: func(cmd *cobra.Command, args []string) error {
		return instanceUsecase.List(cmd.Context())
	},
}

// getFirewallCmd represents the firewall command
var getFirewallCmd = &cobra.Command{
	Use:     "firewall",
	Aliases: []string{"fw", "firewalls"},
	Short:   "Get firewall(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			err := firewallUsecase.Get(cmd.Context(), args[0])
			if err != nil {
				return err
			}
		} else {
			err := firewallUsecase.List(cmd.Context())
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	getCmd.AddCommand(getInstanceCmd)
	getCmd.AddCommand(getFirewallCmd)
	rootCmd.AddCommand(getCmd)
}
