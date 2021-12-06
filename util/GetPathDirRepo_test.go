package util_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetPathDirRepo(t *testing.T) {
	// Same as t.TempDir but for Go 1.14 compatibility
	pathDirRepo, cleanup := util.GetTempDir()
	defer cleanup()

	// Create dummy directory structure
	//   .
	// 	 ├── .git/
	// 	 └── foo/
	// 	     └── bar/
	// 	         └── buzz/
	pathDirGit := filepath.Join(pathDirRepo, ".git")
	pathDirDeep := filepath.Join(pathDirRepo, "foo", "bar", "buzz")

	if err := os.Mkdir(pathDirGit, 0o777); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll(pathDirDeep, 0o777); err != nil {
		t.Fatal(err)
	}

	// Change dir to foo/bar/buzz/
	goBackOrign := util.ChDir(pathDirDeep)
	defer goBackOrign()

	expect := pathDirRepo
	actual := util.GetPathDirRepo()

	assert.Equal(t, expect, actual)
}

func TestGetPathDirRepo_fail_getwd(t *testing.T) {
	oldOsGetwd := util.OsGetwd
	defer func() {
		util.OsGetwd = oldOsGetwd
	}()

	// Mock OsGetwd
	util.OsGetwd = func() (dir string, err error) {
		return "", errors.New("dummy error")
	}

	result := util.GetPathDirRepo()

	assert.Empty(t, result, "on os.Getwd error it should return empty")
}

func TestGetPathDirRepo_up_to_root(t *testing.T) {
	oldOsGetwd := util.OsGetwd
	defer func() {
		util.OsGetwd = oldOsGetwd
	}()

	// Mock OsGetwd
	util.OsGetwd = func() (dir string, err error) {
		return "/", nil
	}

	result := util.GetPathDirRepo()

	assert.Empty(t, result, "searching up to directory root should return empty")
}
