package util_test

import (
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

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
	err := xerrors.New("foo")

	// Test
	out := capturer.CaptureStderr(func() {
		util.ExitOnErr(err)
	})

	// Assertions
	assert.Equal(t, 1, capStatus, "the exit status should be 1 on error")
	assert.Contains(t, out, "ExitOnError_test.go:27", "the error should contain the line number")
	assert.Contains(t, out, "foo", "the stderr should contain the original error message")
}
