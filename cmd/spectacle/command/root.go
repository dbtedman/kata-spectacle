package command

import (
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
	"github.com/spf13/cobra"
)

type Root struct {
	GitLab gitlab.GitLab
}

func (receiver Root) Cobra() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "spectacle",
		Short: "Discover projects within a hosted git platform that contain an OpenAPI Specifications so that an index can be generated.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	serveCommand := Serve{}
	discoverCommand := Discover{
		GitLab: receiver.GitLab,
	}

	rootCommand.AddCommand(serveCommand.Cobra())
	rootCommand.AddCommand(discoverCommand.Cobra())

	return rootCommand
}
