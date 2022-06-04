package command_test

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscoverCobra(t *testing.T) {
	discoverCommand := command.Discover{}

	assert.NotNil(t, discoverCommand.Cobra())
}
