package main_test

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootCobra(t *testing.T) {
	rootCommand := main.Root{}

	assert.NotNil(t, rootCommand.Cobra())
}

func TestRootCobraExecute(t *testing.T) {
	rootCommand := main.Root{}
	rootCommandCobra := rootCommand.Cobra()

	assert.Nil(t, rootCommandCobra.Execute())
}
