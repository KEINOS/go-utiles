package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/kami-zh/go-capturer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleExitOnErr() {
	/*
		Example to mock OsExit in ExitOnErr
	*/
	// Backup and defer restoration
	oldOsExit := util.OsExit
	defer func() {
		util.OsExit = oldOsExit
	}()

	// Mock OsExit
	util.OsExit = func(code int) {
		fmt.Println("the exit code was:", code)
	}

	// Create error
	err := errors.New("foo")

	util.ExitOnErr(err)

	// Output:
	// the exit code was: 1
}

// ----------------------------------------------------------------------------
//  Test
// ----------------------------------------------------------------------------

func TestExitOnErr(t *testing.T) {
	// Backup and recover os.Exit
	oldOsExit := util.OsExit
	defer func() {
		util.OsExit = oldOsExit
	}()

	var capStatus int // This should turn to 1

	// Mock os.Exit
	util.OsExit = func(code int) {
		capStatus = code
	}

	// Create error
	err := errors.New("foo")

	// Wrap the error (stack error)
	err = errors.Wrap(err, "wrapped")

	// Test
	out := capturer.CaptureStderr(func() {
		util.ExitOnErr(err)
	})

	// Assertions
	assert.Equal(t, 1, capStatus, "the exit status should be 1 on error")
	assert.Contains(t, out, "ExitOnError_test.go:60", "stacked error should contain the line number")
	assert.Contains(t, out, "ExitOnError_test.go:63", "stacked error should contain the line number")
	assert.Contains(t, out, "foo", "the stderr should contain the original error message")
}
