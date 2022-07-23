package main_test

import (
	"github.com/dbtedman/kata-spectacle/cmd/spectacle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscoverCobra(t *testing.T) {
	discoverCommand := main.Discover{}

	assert.NotNil(t, discoverCommand.Cobra())
}
