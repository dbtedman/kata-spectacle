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
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCommand.AddCommand(Serve{}.Cobra())
	rootCommand.AddCommand(Discover{GitLab: receiver.GitLab}.Cobra())

	return rootCommand
}
