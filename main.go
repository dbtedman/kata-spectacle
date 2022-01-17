package main

import (
	"github.com/dbtedman/kata-spectacle/cmd"
	"log"
)

func main() {
	rootCommand := cmd.Cmd{}
	if err := rootCommand.Listen(); err != nil {
		log.Fatal(err)
	}
}
