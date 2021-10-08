package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleIsDir() {
	// Existing dir
	if util.IsDir("../testdata/sample_data") {
		fmt.Println("is dir")
	}

	// Not existing dir
	if !util.IsDir("./foobar") {
		fmt.Println("not a dir")
	}

	// File exists but not a dir
	if !util.IsDir("./IsDir_test.go") {
		fmt.Println("not a dir")
	}

	// Output:
	// is dir
	// not a dir
	// not a dir
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestIsDir(t *testing.T) {
	for path, expect := range map[string]bool{
		"sample_data":             true,
		"sample_data/sample.json": false,
		"sample_data/.sample":     false,
		"unknown-dir":             false,
		"unknown-file":            false,
		".unknown-dotfile":        false,
	} {
		actual := util.IsDir("../testdata/" + path)
		assert.Equal(t, expect, actual, "Failed data: %s", path)
	}
}
