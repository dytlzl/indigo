package cmd

import (
	"github.com/dytlzl/indigo/cmd/di"
	"github.com/dytlzl/indigo/pkg/config"
	"github.com/spf13/cobra"
)

func NewStopCmd(conf config.Config) *cobra.Command {
	var force bool
	cmd := &cobra.Command{
		Use:   "stop [name]",
		Short: "Stop an instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if force {
				return di.InitializeInstanceUseCase(conf).ForceStop(cmd.Context(), args[0])
			} else {
				return di.InitializeInstanceUseCase(conf).Stop(cmd.Context(), args[0])
			}
		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "indigo stop node00 --force")
	return cmd
}
