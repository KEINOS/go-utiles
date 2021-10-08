package util

import (
	"os"
)

// IsDir returns true if pathFile is an existing directory and not a file.
func IsDir(pathFile string) bool {
	fileInfo, err := os.Stat(pathFile)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
