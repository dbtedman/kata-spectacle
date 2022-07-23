package main_test

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServeCobra(t *testing.T) {
	serveCommand := main.Serve{}

	assert.NotNil(t, serveCommand.Cobra())
}

func TestServeCobraExecute(t *testing.T) {
	serveCommand := main.Serve{}
	serveCommandCobra := serveCommand.Cobra()

	assert.Nil(t, serveCommandCobra.Execute())
}
