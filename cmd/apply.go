package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a manifest file",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			return err
		}
		fileBody, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		manifestFile := ManifestFile{}
		err = yaml.Unmarshal(fileBody, &manifestFile)
		if err != nil {
			return err
		}
		switch manifestFile.Kind {
		case "Firewall":
			err = firewallUsecase.Apply(cmd.Context(), fileBody)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid kind was specified: %s", manifestFile.Kind)
		}
		return nil
	},
}

type ManifestFile struct {
	Kind string `yaml:"kind"`
}

func init() {
	applyCmd.Flags().StringP("filename", "f", "[]", "indigo -f instance.yaml")
	applyCmd.MarkFlagRequired("filename")
	rootCmd.AddCommand(applyCmd)
}
