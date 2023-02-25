package cmd

import (
	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {
	var getInstanceCmd = &cobra.Command{
		Use:     "instance",
		Aliases: []string{"i", "instances"},
		Short:   "Get instance(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return instanceUseCase.List(cmd.Context())
		},
	}
	var getOSCmd = &cobra.Command{
		Use:     "os",
		Aliases: []string{"oses"},
		Short:   "Get OS(es)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return osUseCase.List(cmd.Context())
		},
	}

	var getSSHKeyCmd = &cobra.Command{
		Use:     "sshkey",
		Aliases: []string{"sk", "sshkeys"},
		Short:   "Get SSH Key(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sshKeyUseCase.List(cmd.Context())
		},
	}

	var getPlanCmd = &cobra.Command{
		Use:     "plan",
		Aliases: []string{"p", "plans"},
		Short:   "Get plan(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return planUseCase.List(cmd.Context())
		},
	}

	// getFirewallCmd represents the firewall command
	var getFirewallCmd = &cobra.Command{
		Use:     "firewall",
		Aliases: []string{"fw", "firewalls"},
		Short:   "Get firewall(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				err := firewallUseCase.Get(cmd.Context(), args[0])
				if err != nil {
					return err
				}
			} else {
				err := firewallUseCase.List(cmd.Context())
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get resource(s)",
	}
	cmd.AddCommand(
		getInstanceCmd,
		getFirewallCmd,
		getOSCmd,
		getSSHKeyCmd,
		getPlanCmd,
	)
	return cmd
}
