package test

import (
	"testing"

	"github.com/awgst/goes/generator"
	"github.com/stretchr/testify/assert"
)

func TestRequestGeneratorMakeFunctionWillGenerateFile(t *testing.T) {
	var requestGenerator generator.Request
	tmpDir := t.TempDir()

	assert.Nil(t, requestGenerator.Make("Request", tmpDir, "request"))
}
