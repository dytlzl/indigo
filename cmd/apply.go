package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func NewApplyCmd() *cobra.Command {
	var (
		filename string
	)
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a manifest file",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileBody, err := os.ReadFile(filename)
			if err != nil {
				return err
			}
			manifestFile := struct {
				Kind string `yaml:"kind"`
			}{}
			err = yaml.Unmarshal(fileBody, &manifestFile)
			if err != nil {
				return err
			}
			switch manifestFile.Kind {
			case "Firewall":
				err = firewallUseCase.Apply(cmd.Context(), fileBody)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("invalid kind was specified: %s", manifestFile.Kind)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&filename, "filename", "f", "[]", "indigo -f instance.yaml")
	err := cmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalln(err)
	}
	return cmd
}
