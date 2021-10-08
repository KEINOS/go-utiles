package util

import (
	"os"
	"path/filepath"
)

// WriteTmpFile saves the string of data to a temp file. It returns the saved path
// and a function to delete that temp file.
func WriteTmpFile(data string) (pathSaved string, deferCleanUp func(), err error) {
	nameFile := "tmp" + RandStr(32)
	pathFile := filepath.Join(os.TempDir(), nameFile)

	// Create and write data (Go v1.14 compatible)
	pFile, err := os.Create(pathFile)
	if err == nil {
		defer func() { pFile.Close() }()

		pathSaved = pFile.Name()
		deferCleanUp = func() {
			os.Remove(pathSaved) // clean up
		}
		_, err = pFile.WriteString(data)
	}

	return pathSaved, deferCleanUp, err
}
