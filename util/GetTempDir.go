package util

import (
	"os"
	"path/filepath"
)

// GetTempDir returns a temporary directory and the cleanup function for the test
// to use. It is similar to T.TempDir() but for Go 1.14 compatibility.
func GetTempDir() (pathDir string, cleanup func()) {
	pathDirTemp := os.TempDir()
	nameDirNew := RandStr(32)
	pathDirNew := filepath.Join(pathDirTemp, nameDirNew)

	err := os.MkdirAll(pathDirNew, 0o766)
	ExitOnErr(err)

	return pathDirNew, func() {
		os.RemoveAll(pathDirNew)
	}
}
