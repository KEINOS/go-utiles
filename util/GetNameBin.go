package util

import (
	"os"
	"path"
)

// GetNameBin returns the file name of the current executable binary.
func GetNameBin() string {
	result := "unknown"

	if os.Args[0] != "" {
		result = path.Base(os.Args[0])
	}

	if exe, err := os.Executable(); err == nil {
		result = path.Base(exe)
	}

	return result
}
