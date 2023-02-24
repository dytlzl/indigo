package cmd

import (
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop [name]",
	Short: "Stop an instance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, err := cmd.PersistentFlags().GetBool("force")
		if err != nil {
			return err
		}
		if force {
			err = instanceUseCase.ForceStop(cmd.Context(), args[0])
			if err != nil {
				return err
			}
		} else {
			err = instanceUseCase.Stop(cmd.Context(), args[0])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	stopCmd.PersistentFlags().BoolP("force", "f", false, "indigo stop node00 --force")
	rootCmd.AddCommand(stopCmd)
}
