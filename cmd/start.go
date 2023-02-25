package cmd

import (
	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [name]",
		Short: "Start an instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return instanceUseCase.Start(cmd.Context(), args[0])
		},
	}
	return cmd
}
