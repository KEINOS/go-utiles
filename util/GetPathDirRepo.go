package util

import (
	"path/filepath"
)

// GetPathDirRepo returns the root directory of the current git repo.
// If no ".git" directory found then returns "".
//
// It will search up the directory from the current working dir upto the depth level.
func GetPathDirRepo() string {
	pathDirCurr, err := OsGetwd()
	if err != nil {
		return ""
	}

	// Search current dir
	path := pathDirCurr

	// Search up parent dirs
	for {
		if IsDir(filepath.Join(path, ".git")) {
			return path
		}

		path = filepath.Dir(path) // Step up to parent dir

		if path == "." || path == "/" {
			break
		}
	}

	return ""
}
