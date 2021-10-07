package util

import "os"

// PathExists returns true if the path is an existing file or dir.
func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
