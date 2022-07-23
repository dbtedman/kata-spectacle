package main

import (
	"github.com/spf13/cobra"
)

type Serve struct{}

func (receiver Serve) Cobra() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {},
	}
}
