package test

import (
	"testing"

	"github.com/awgst/goes/generator"
	"github.com/stretchr/testify/assert"
)

func TestMakeFunctionWillGenerateFil(t *testing.T) {
	var controller generator.Controller
	tmpDir := t.TempDir()

	assert.Nil(t, controller.Make("Controller", tmpDir, "controller"))
}
