package util

import (
	"os"
)

var (
	// OsGetwd is a copy for os.Getwd to ease mocking during test.
	//
	// All functions of this package that needs to use os.Getwd uses OsGetwd instead.
	// See the example in the test of ChDir for how-to-mock.
	OsGetwd = os.Getwd
	// OsChdir is a copy for os.Chdir to ease mocking during test.
	//
	// All functions of this package that needs to use os.Chdir uses OsChdir instead.
	// See the example in the test of ChDir for how-to-mock.
	OsChdir = os.Chdir
)

// ChDir changes the current working directory to the given path in one-go.
// It returns a function which moves back to the original directory.
//
// Note: This function exits with status 1 if any error happens.
func ChDir(pathDir string) (deferReturn func()) {
	// Get current working dir before move (OsGetwd is a copy of os.Getwd)
	oldPwd, err := OsGetwd()
	ExitOnErr(err)

	// Change dir (OsChdir is a copy of os.Chdir)
	err = OsChdir(pathDir)
	ExitOnErr(err)

	return func() {
		err = os.Chdir(oldPwd)
	}
}
