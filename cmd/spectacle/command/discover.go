package command

import (
	"fmt"
	"github.com/dbtedman/kata-spectacle/internal/domain/discover"
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
	"github.com/spf13/cobra"
	"log"
)

type Discover struct {
	GitLab gitlab.GitLab
}

func (receiver Discover) Cobra() *cobra.Command {
	discovery := discover.Discover{
		GitLab: receiver.GitLab,
	}

	return &cobra.Command{
		Use: "discover",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := discovery.Execute(discover.DiscoverRequest{})

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
