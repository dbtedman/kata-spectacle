package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Serve struct {
}

func (receiver Serve) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("")
		},
	}
}
