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
	rootCommand := command.Root{
		GitLab: gitlab.GitLab{
			Token: os.Getenv(command.GitLabTokenEnv),
			Url:   os.Getenv(command.GitLabURLEnv),
		},
	}

	return rootCommand.Cobra().Execute()
}
