package cmd

import (
	"github.com/dbtedman/kata-spectacle/internal/gateway"
	"github.com/spf13/cobra"
	"os"
)

type Cmd struct {
}

func (receiver Cmd) Listen() error {
	var rootCmd = &cobra.Command{
		Use:   "spectacle",
		Short: "An auto discovery api index",
		Long:  "An auto discovery api index",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	gitlabToken := os.Getenv("APIS_GITLAB_TOKEN")
	gitlabUrl := os.Getenv("APIS_GITLAB_URL")

	gitlab := gateway.GitLab{Token: gitlabToken, Url: gitlabUrl}

	serveCmd := Serve{}
	discoverCmd := Discover{
		GitLab: gitlab,
	}

	rootCmd.AddCommand(serveCmd.Cmd())
	rootCmd.AddCommand(discoverCmd.Cmd())

	return rootCmd.Execute()
}
