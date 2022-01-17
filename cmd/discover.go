package cmd

import (
	"fmt"
	"github.com/dbtedman/kata-spectacle/internal/domain"
	"github.com/dbtedman/kata-spectacle/internal/gateway"
	"github.com/spf13/cobra"
	"log"
)

type Discover struct {
	GitLab gateway.GitLab
}

func (receiver Discover) Cmd() *cobra.Command {
	discover := domain.Discover{
		GitLab: receiver.GitLab,
	}

	return &cobra.Command{
		Use:   "discover",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := discover.Execute(domain.DiscoverRequest{})

			if err != nil {
				log.Fatalln(err)
			}

			// TODO: Would we now write this to file system now?
			for _, result := range results {
				fmt.Printf("%+v\n", result)
			}
		},
	}
}
