package command_test

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServeCobra(t *testing.T) {
	serveCommand := command.Serve{}

	assert.NotNil(t, serveCommand.Cobra())
}

func TestServeCobraExecute(t *testing.T) {
	serveCommand := command.Serve{}
	serveCommandCobra := serveCommand.Cobra()

	assert.Nil(t, serveCommandCobra.Execute())
}
