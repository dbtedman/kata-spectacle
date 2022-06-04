package command_test

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootCobra(t *testing.T) {
	rootCommand := command.Root{}

	assert.NotNil(t, rootCommand.Cobra())
}

func TestRootCobraExecute(t *testing.T) {
	rootCommand := command.Root{}
	rootCommandCobra := rootCommand.Cobra()

	assert.Nil(t, rootCommandCobra.Execute())
}
