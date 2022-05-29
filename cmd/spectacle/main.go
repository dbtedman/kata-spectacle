package main

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle/action"
	"log"
)

func main() {
	root := action.Root{}
	if err := root.Listen(); err != nil {
		log.Fatal(err)
	}
}
