package util

import "os"

// ChDirHome is similar to util.ChDir but it moves the current working directory
// to the user's home directory in one-go.
// It returns a function to move back to the original directory.
//
// Note: This function exits with status 1 if any error happens.
func ChDirHome() func() {
	pathDirHome, err := os.UserHomeDir()
	ExitOnErr(err)

	return ChDir(pathDirHome)
}
