package util

import (
	"fmt"
	"os"
)

// OsExit is a copy of os.Exit to ease mocking during test.
//
// All functions of this package that needs to use os.Exit uses OsExit instead.
// See the example of ExitOnError for how-to-mock.
var OsExit = os.Exit

// ExitOnErr exits with status 1 if err is not nil.
//
// To test this function, mock the OsExit function variable.
// See ExitOnError_test.go for an example.
func ExitOnErr(err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%+v\n", err)

	OsExit(1)
}
