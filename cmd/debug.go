package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use: "debug",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := client.Get(cmd.Context(), "/vm/sshkey")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	},
}
