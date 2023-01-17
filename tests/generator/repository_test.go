package test

import (
	"testing"

	"github.com/awgst/goes/generator"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGeneratorMakeFunctionWillGenerateFile(t *testing.T) {
	var repository generator.Repository
	tmpDir := t.TempDir()

	assert.Nil(t, repository.Make("Repository", tmpDir, "repository"))
}
