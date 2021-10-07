package util_test

import (
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
)

func TestPathExists(t *testing.T) {
	for path, expect := range map[string]bool{
		"sample_data":             true,
		"sample_data/sample.json": true,
		"sample_data/.sample":     true,
		"unknown-dir":             false,
		"unknown-file":            false,
		".unknown-dotfile":        false,
	} {
		actual := util.PathExists("../testdata/" + path)
		assert.Equal(t, expect, actual, "Failed data: %s", path)
	}
}
