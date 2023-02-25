package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version  string = "not provided"
	Revision string = "not provided"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:    "version",
	Short:  "Show version",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Client Version: %s, Commit: %s\n", Version, Revision)
	},
}
