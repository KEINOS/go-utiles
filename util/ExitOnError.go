package util

import (
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

var OsExit = os.Exit

// ExitOnErr exits with status 1 if err is not nil.
func ExitOnErr(err error) {
	if err != nil {
		wrap := xerrors.Errorf("Errors from: %+w", err)
		fmt.Fprintf(os.Stderr, "%v", wrap)

		OsExit(1)
	}
}
