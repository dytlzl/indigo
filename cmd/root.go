package cmd

import (
	"context"
	"os/user"
	"path/filepath"
	"time"

	"github.com/dytlzl/indigo/pkg/config"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := &cobra.Command{
		Use:   "indigo",
		Short: "Indigo API client",
		Long:  "indigo is a Indigo API client written in Go.",
	}
	var configFilename string
	u, err := user.Current()
	if err != nil {
		return err
	}
	cmd.PersistentFlags().StringVar(&configFilename, "config", filepath.Join(u.HomeDir, ".indigo.yaml"), "config file (default is $HOME/.indigo.yaml)")
	conf := config.NewConfig(configFilename)
	cmd.AddCommand(
		NewApplyCmd(conf),
		NewCreatCmd(conf),
		NewDeleteCmd(conf),
		NewGetCmd(conf),
		NewStartCmd(conf),
		NewStopCmd(conf),
		NewDebugCmd(conf),
		versionCmd,
	)
	return cmd.ExecuteContext(ctx)
}
