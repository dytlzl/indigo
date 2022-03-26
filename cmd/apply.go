package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/repository"
	"github.com/dytlzl/indigo/pkg/usecase"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a manifest file",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			log.Fatalln(err)
		}
		fileBody, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}
		manifestFile := ManifestFile{}
		err = yaml.Unmarshal(fileBody, &manifestFile)
		if err != nil {
			log.Fatalln(err)
		}
		client, err := api.NewClient(conf)
		if err != nil {
			log.Fatalln(err)
		}
		switch manifestFile.Kind {
		case "Firewall":
			fr := repository.NewAPIFirewallRepository(client)
			ir := repository.NewAPIInstanceRepository(client)
			u := usecase.NewFirewallUsecase(fr, ir)
			err = u.Apply(ctx, fileBody)
			if err != nil {
				log.Fatalln(err)
			}
		}
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
