package cmd

import (
	"github.com/dytlzl/indigo/cmd/di"
	"github.com/dytlzl/indigo/pkg/config"
	"github.com/spf13/cobra"
)

func NewStartCmd(conf config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [name]",
		Short: "Start an instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return di.InitializeInstanceUseCase(conf).Start(cmd.Context(), args[0])
		},
	}
	return cmd
}
