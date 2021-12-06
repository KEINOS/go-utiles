package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleGetTempDir() {
	pathDir, cleanup := util.GetTempDir()
	defer cleanup()

	if util.IsDir(pathDir) {
		fmt.Println("directory exists")
	}
	// Output: directory exists
}

func TestGetTempDir(t *testing.T) {
	path, deferFunc := util.GetTempDir()
	require.True(t, util.IsDir(path))

	// Crete dummy file
	f, err := util.CreateTemp(path, "*")
	require.NoError(t, err)

	pathFile := f.Name()
	f.Close()

	assert.True(t, util.IsFile(pathFile))

	// Cleanup
	deferFunc()

	assert.False(t, util.IsDir(path))
	assert.False(t, util.IsFile(pathFile))
}
