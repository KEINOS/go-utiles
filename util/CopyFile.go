package util

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

// OsOpen is a copy of os.Open to ease mock during test.
var OsOpen = os.Open

// OsCreate is a copy of os.Create to ease mock during test.
var OsCreate = os.Create

// IoCopy is a copy of io.Copy to ease mock during test.
var IoCopy = io.Copy

// CopyFile copies the file from to.
func CopyFile(from, to string) error {
	if !IsFile(from) {
		return errors.Errorf("failed to copy. input path is not a file: %v", from)
	}

	if IsDir(to) {
		return errors.Errorf("failed to copy. output path is a dir: %v", to)
	}

	// Open original file
	in, err := OsOpen(from)
	if err != nil {
		return errors.Wrap(err, "failed to open file: "+from)
	}
	defer in.Close()

	// Create new file
	out, err := OsCreate(to)
	if err != nil {
		return errors.Wrap(err, "failed to create file before write: "+to)
	}
	defer out.Close()

	// Write data
	if _, err := IoCopy(out, in); err != nil {
		return errors.Wrap(err, "failed to write file: "+to)
	}

	return nil
}
