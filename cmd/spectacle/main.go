package main

import (
	"fmt"
	"github.com/dbtedman/kata-spectacle/internal/domain/entity"
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
	"os"
)

func main() {
	if err := executeCommand(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func executeCommand() error {
	// TODO: modify this behaviour to leverage https://github.com/spf13/viper for config
	// config via: cli args, environment variables, config file, config server (e.g. vault)
	// consider how app can respond to config settings being changed while the app is live

	config := entity.Config{
		Token: os.Getenv(GitLabTokenEnv),
		Url:   os.Getenv(GitLabURLEnv),
	}

	// rather than providing the config directly, do we define a function that can be called
	// to access the current config value?
	rootCommand := Root{
		GitLab: gitlab.GitLab{
			Config: config,
		},
	}

	return rootCommand.Cobra().Execute()
}
