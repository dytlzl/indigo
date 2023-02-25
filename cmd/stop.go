package cmd

import (
	"github.com/spf13/cobra"
)

func NewStopCmd() *cobra.Command {
	var force bool
	cmd := &cobra.Command{
		Use:   "stop [name]",
		Short: "Stop an instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if force {
				err := instanceUseCase.ForceStop(cmd.Context(), args[0])
				if err != nil {
					return err
				}
			} else {
				err := instanceUseCase.Stop(cmd.Context(), args[0])
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "indigo stop node00 --force")
	return cmd
}
