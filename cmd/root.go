package cmd

import (
	"log"
	"os"
	"os/user"

	"github.com/dytlzl/indigo/pkg/config"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "indigo",
	Short: "Indigo API client",
	Long:  "indigo is a Indigo API client written in Go.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var configFile string

var conf config.Config

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	rootCmd.PersistentFlags().StringVar(&configFile, "config", u.HomeDir+"/.indigo.yaml", "config file (default is $HOME/.indigo.yaml)")
	cobra.OnInitialize(func() {
		conf = config.NewConfig(configFile)
	})
}
