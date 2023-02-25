package cmd

import (
	"github.com/dytlzl/indigo/cmd/di"
	"github.com/dytlzl/indigo/pkg/config"
	"github.com/dytlzl/indigo/pkg/infra/cmdutil"
	"github.com/spf13/cobra"
)

func NewCreatCmd(conf config.Config) *cobra.Command {
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
			return di.InitializeInstanceUseCase(conf).Create(cmd.Context(), args[0], planID, osID, regionID, sshKeyID)
		},
	}
	cmd.Flags().AddFlagSet(cmdutil.FlagSetBuilder().
		Int(&planID, "plan-id", "", -1, "plan(Size) ID", cobra.MarkFlagRequired).
		Int(&sshKeyID, "ssh-key-id", "", -1, "SSH key ID", cobra.MarkFlagRequired).
		Int(&osID, "os-id", "", 13, "OS ID (Ubuntu 22.04 13)").
		Int(&regionID, "region-id", "", 1, "region ID").
		Build(),
	)
	cmd.AddCommand(createInstanceCmd)
	return cmd
}
