package util_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/kami-zh/go-capturer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleChDir() {
	pathDirToMove := "/tmp"

	// Move working directory and defer switch back the diretory
	funcReturn := util.ChDir(pathDirToMove)
	defer funcReturn()

	pathDirCurrent, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pathDirCurrent)
	// Output:
	// /tmp
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestChDir_fail_get_working_dir(t *testing.T) {
	// Backup and defer restoration util.OsGetwd
	oldOsGetwd := util.OsGetwd
	defer func() {
		util.OsGetwd = oldOsGetwd
	}()

	// Backup and defer restoration util.OsExit
	oldOsExit := util.OsExit
	defer func() {
		util.OsExit = oldOsExit
	}()

	// Mock the os.Getwd in util.OsGetwd
	util.OsGetwd = func() (dir string, err error) {
		return "", errors.New("foo\n") // Error occurred in line 55
	}

	// Mock the os.Getwd in util.OsGetwd
	captureStatus := 0 // it should turn into 1 on fail
	util.OsExit = func(code int) {
		captureStatus = code
	}

	// Test
	out := capturer.CaptureStderr(func() {
		_ = util.ChDir("/foobarbuz")
	})

	// Assertion
	assert.Equal(t, 1, captureStatus, "on error it shuld end with status 1")
	assert.Contains(t, out, "ChDir_test.go:55", "it should contain the error line")
	assert.Contains(t, out, "no such file or directory", "it should contain the error reason")
}
