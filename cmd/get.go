package cmd

import (
	"github.com/dytlzl/indigo/cmd/di"
	"github.com/dytlzl/indigo/pkg/config"
	"github.com/spf13/cobra"
)

func NewGetCmd(conf config.Config) *cobra.Command {
	var getInstanceCmd = &cobra.Command{
		Use:     "instance",
		Aliases: []string{"i", "instances"},
		Short:   "Get instance(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return di.InitializeInstanceUseCase(conf).List(cmd.Context())
		},
	}
	var getOSCmd = &cobra.Command{
		Use:     "os",
		Aliases: []string{"oses"},
		Short:   "Get OS(es)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return di.InitializeOSUseCase(conf).List(cmd.Context())
		},
	}
	var getSSHKeyCmd = &cobra.Command{
		Use:     "sshkey",
		Aliases: []string{"sk", "sshkeys"},
		Short:   "Get SSH Key(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return di.InitializeSSHKeyUseCase(conf).List(cmd.Context())
		},
	}
	var getPlanCmd = &cobra.Command{
		Use:     "plan",
		Aliases: []string{"p", "plans"},
		Short:   "Get plan(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return di.InitializePlanUseCase(conf).List(cmd.Context())
		},
	}
	var getFirewallCmd = &cobra.Command{
		Use:     "firewall",
		Aliases: []string{"fw", "firewalls"},
		Short:   "Get firewall(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return di.InitializeFirewallUseCase(conf).Get(cmd.Context(), args[0])
			} else {
				return di.InitializeFirewallUseCase(conf).List(cmd.Context())
			}
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
