//nolint: dupl // let duplicate lines
package util_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Examples
// ----------------------------------------------------------------------------

func ExampleCopyFile() {
	// Get file paths to copy from and to
	from := getPathFileToCopyFrom()
	to := getPathFileToCopyTo()

	// Copy file
	if err := util.CopyFile(from, to); err != nil {
		log.Fatal(err)
	}

	// Read copied contents for testing
	contByte, err := os.ReadFile(to)
	if err != nil {
		log.Fatal(err)
	}

	// Require Equal
	expect := "This is a dummy text at /testdata/dummy.txt.\n"

	if actual := string(contByte); expect != actual {
		log.Fatalf("output file does not match to input.\nexpect: %vactual: %v", expect, actual)
	}

	fmt.Println("ok")

	// Clean up copied file
	if err = os.RemoveAll(to); err != nil {
		log.Fatal(err)
	}

	// Output: ok
}

func getPathFileToCopyFrom() string {
	return filepath.Join(util.GetPathDirRepo(), "testdata", "dummy.txt")
}

func getPathFileToCopyTo() string {
	pathDirTemp := os.TempDir()
	nameFileTemp := util.RandStr(32) + ".txt"

	return filepath.Join(pathDirTemp, nameFileTemp)
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestCopyFile_path_is_a_dir(t *testing.T) {
	pathDirTemp, cleanup := util.GetTempDir()
	defer cleanup()

	// Create
	fp, err := util.CreateTemp(pathDirTemp, "*")
	require.NoError(t, err, "failed to open temp file")

	pathFileTemp := fp.Name()

	err = fp.Close()
	require.NoError(t, err, "failed to close temp file")

	// from is a dir
	err = util.CopyFile(pathDirTemp, pathFileTemp)
	require.Error(t, err, "it should return an error if the path is a dir")
	assert.Contains(t, err.Error(), "failed to copy. input path is not a file")

	// to is a dir
	err = util.CopyFile(pathFileTemp, pathDirTemp)
	require.Error(t, err, "it should return an error if the path is a dir")
	assert.Contains(t, err.Error(), "failed to copy. output path is a dir")
}

func TestCopyFile_fail_to_open_file(t *testing.T) {
	// Backup and defer recover
	oldOsOpen := util.OsOpen
	defer func() {
		util.OsOpen = oldOsOpen
	}()

	// Mock
	util.OsOpen = func(name string) (*os.File, error) {
		return nil, errors.New("forced error")
	}

	// Create dummy file
	pathDirTemp, cleanup := util.GetTempDir()
	defer cleanup()

	fp, err := util.CreateTemp(pathDirTemp, "*")
	require.NoError(t, err, "failed to open temp file")

	pathFileTemp := fp.Name()

	err = fp.Close()
	require.NoError(t, err, "failed to close temp file")

	// Assertion
	err = util.CopyFile(pathFileTemp, pathFileTemp)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open file")
}

func TestCopyFile_fail_to_create_file(t *testing.T) {
	// Backup and defer recover
	oldOsCreate := util.OsCreate
	defer func() {
		util.OsCreate = oldOsCreate
	}()

	// Mock
	util.OsCreate = func(name string) (*os.File, error) {
		return nil, errors.New("forced error")
	}

	// Create dummy file
	pathDirTemp, cleanup := util.GetTempDir()
	defer cleanup()

	fp, err := util.CreateTemp(pathDirTemp, "*")
	require.NoError(t, err, "failed to open temp file")

	pathFileTemp := fp.Name()

	err = fp.Close()
	require.NoError(t, err, "failed to close temp file")

	// Assertion
	err = util.CopyFile(pathFileTemp, pathFileTemp)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create file before write")
}

func TestCopyFile_fail_to_read_written_file(t *testing.T) {
	// Backup and defer recover
	oldOsCreate := util.OsCreate
	defer func() {
		util.OsCreate = oldOsCreate
	}()

	// Mock
	util.OsCreate = func(name string) (*os.File, error) {
		return nil, nil
	}

	// Create dummy file
	pathDirTemp, cleanup := util.GetTempDir()
	defer cleanup()

	fp, err := util.CreateTemp(pathDirTemp, "*")
	require.NoError(t, err, "failed to open temp file")

	pathFileTemp := fp.Name()

	err = fp.Close()
	require.NoError(t, err, "failed to close temp file")

	// Assertion
	err = util.CopyFile(pathFileTemp, pathFileTemp)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to write file")
}
