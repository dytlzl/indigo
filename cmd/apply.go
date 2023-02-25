package cmd

import (
	"fmt"
	"os"

	"github.com/dytlzl/indigo/cmd/di"
	"github.com/dytlzl/indigo/pkg/config"
	"github.com/dytlzl/indigo/pkg/infra/cmdutil"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func NewApplyCmd(conf config.Config) *cobra.Command {
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
				return di.InitializeFirewallUseCase(conf).Apply(cmd.Context(), fileBody)
			default:
				return fmt.Errorf("invalid kind was specified: %s", manifestFile.Kind)
			}
		},
	}
	cmd.Flags().AddFlagSet(cmdutil.FlagSetBuilder().
		String(&filename, "filename", "f", "[]", "Path to the file that contains the configuration to apply", cobra.MarkFlagRequired, cmdutil.MarkFlagFilename("yaml", "yml")).
		Build(),
	)
	return cmd
}
