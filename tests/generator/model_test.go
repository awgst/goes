package test

import (
	"testing"

	"github.com/awgst/goes/generator"
	"github.com/stretchr/testify/assert"
)

func TestMakeFunctionWillGenerateFile(t *testing.T) {
	var modelGenerator generator.Model
	tmpDir := t.TempDir()

	assert.Nil(t, modelGenerator.Make("Model", tmpDir, "model"))
}
