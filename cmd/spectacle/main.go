package main

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle/command"
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
	"log"
	"os"
)

func main() {
	if err := executeCommand(); err != nil {
		log.Fatal(err)
	}
}

func executeCommand() error {
	gitlabToken := os.Getenv("APIS_GITLAB_TOKEN")
	gitlabUrl := os.Getenv("APIS_GITLAB_URL")

	rootCommand := command.Root{
		GitLab: gitlab.GitLab{
			Token: gitlabToken,
			Url:   gitlabUrl,
		},
	}

	return rootCommand.Cobra().Execute()
}
