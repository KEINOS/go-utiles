package util

import (
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

// OsExit is a copy for os.Exit to ease mocking during test.
//
// All functions of this package that needs to os.Exit uses OsExit instead.
// See the example of ExitOnError for how-to-mock.
var OsExit = os.Exit

// ExitOnErr exits with status 1 if err is not nil.
//
// To test this function, mock the OsExit function variable.
// See ExitOnError_test.go for an example.
func ExitOnErr(err error) {
	if err != nil {
		wrap := xerrors.Errorf("Errors from: %+w", err)
		fmt.Fprintf(os.Stderr, "%v", wrap)

		OsExit(1)
	}
}
