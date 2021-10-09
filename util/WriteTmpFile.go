package util

import (
	"os"
)

// WriteTmpFile saves the string of data to a temp file. It returns the saved path
// and a function to delete that temp file.
func WriteTmpFile(data string) (pathSaved string, funcCleanUp func(), err error) {
	pFile, err := CreateTemp("", "tmp-")
	if err == nil {
		defer func() { pFile.Close() }()

		pathSaved = pFile.Name()
		funcCleanUp = func() {
			os.Remove(pathSaved) // clean up
		}
		_, err = pFile.WriteString(data)
	}

	return pathSaved, funcCleanUp, err
}
