package test

import (
	"testing"

	"github.com/awgst/goes/generator"
	"github.com/stretchr/testify/assert"
)

func TestResponseGeneratorMakeFunctionWillGenerateFile(t *testing.T) {
	var responseGenerator generator.Response
	tmpDir := t.TempDir()

	assert.Nil(t, responseGenerator.Make("Response", tmpDir, "response"))
}
