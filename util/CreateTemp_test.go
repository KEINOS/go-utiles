package util_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleCreateTemp() {
	// Create temp file under temp dir with the name "foo-*.json"
	p, err := util.CreateTemp("", "foo-*.json")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer p.Close() // Don't forget to close it

	pathSaved := p.Name() // Get the file path

	// Do  something with the file
	if util.IsFile(pathSaved) {
		fmt.Println("file exists")
	}

	// Clean up the temp file
	os.Remove(pathSaved)

	if !util.IsFile(pathSaved) {
		fmt.Println("temp file cleaned")
	}

	// Output:
	// file exists
	// temp file cleaned
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestCreateTemp_no_asterisk(t *testing.T) {
	p, err := util.CreateTemp("", "foo.json")
	require.NoError(t, err, "non-existing directory should return an error")

	defer p.Close()

	filePath := p.Name()

	assert.FileExists(t, filePath, "if no error the file should exist")
	assert.Contains(t, filePath, "foo", "file name should contain the pattern string")
	assert.Contains(t, filePath, ".json", "file name should contain the extension of the pattern")
	assert.False(t, strings.HasSuffix(filePath, ".json"), "file name should end with random string")
}

func TestCreateTemp_non_existing_dir(t *testing.T) {
	p, err := util.CreateTemp("./foo/bar", "buz-*.json")

	require.Error(t, err, "non-existing directory should return an error")
	assert.Nil(t, p, "on error, filepointer should be nil")
}
