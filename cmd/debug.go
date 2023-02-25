package cmd

import (
	"fmt"

	"github.com/dytlzl/indigo/pkg/config"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/spf13/cobra"
)

func NewDebugCmd(conf config.Config) *cobra.Command {
	return &cobra.Command{
		Use:    "debug",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(conf)
			b, err := client.Get(cmd.Context(), "/vm/sshkey")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
			return nil
		},
	}
}
