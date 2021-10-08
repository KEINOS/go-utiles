package util

import "io/ioutil"

// ReadFile is similar to os.ReadFile inf Go v1.16+.
// Aim to use for Go v1.14 and 1.15 compatibility.
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
