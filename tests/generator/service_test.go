package test

import (
	"testing"

	"github.com/awgst/goes/generator"
	"github.com/stretchr/testify/assert"
)

func TestServiceGeneratorMakeFunctionWillGenerateFile(t *testing.T) {
	var service generator.Service
	tmpDir := t.TempDir()

	assert.Nil(t, service.Make("Service", tmpDir, "service"))
}
