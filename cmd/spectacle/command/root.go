package command

import (
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
	"github.com/spf13/cobra"
)

// GitLabTokenEnv is the environment variable that contains an access token used to
// authenticate api calls to GitLab,
// https://docs.gitlab.com/ee/api/index.html#personalprojectgroup-access-tokens
const GitLabTokenEnv = "SPECTACLE_GITLAB_TOKEN"

// GitLabURLEnv is the environment variable that defines the base URL for API calls to
// the GitLab REST API, https://docs.gitlab.com/ee/api/index.html
const GitLabURLEnv = "SPECTACLE_GITLAB_URL"

type Root struct {
	GitLab gitlab.GitLab
}

func (receiver Root) Cobra() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "spectacle",
		Short: "Discover projects within a hosted git platform that contain an OpenAPI Specifications so that an index can be generated.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCommand.AddCommand(Serve{}.Cobra())
	rootCommand.AddCommand(Discover{GitLab: receiver.GitLab}.Cobra())

	return rootCommand
}
