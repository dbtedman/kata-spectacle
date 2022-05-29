package action

import (
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
	"github.com/spf13/cobra"
	"os"
)

type Root struct {
}

func (receiver Root) Listen() error {
	gitlabToken := os.Getenv("APIS_GITLAB_TOKEN")
	gitlabUrl := os.Getenv("APIS_GITLAB_URL")

	gitlabGateway := gitlab.GitLab{Token: gitlabToken, Url: gitlabUrl}

	serveAction := Serve{}
	discoverAction := Discover{
		GitLab: gitlabGateway,
	}

	var rootCmd = &cobra.Command{
		Use:   "spectacle",
		Short: "An auto discovery api index",
		Long:  "An auto discovery api index",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(serveAction.Cmd())
	rootCmd.AddCommand(discoverAction.Cmd())

	return rootCmd.Execute()
}
