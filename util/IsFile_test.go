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

func ExampleIsFile() {
	fmt.Println(util.IsFile("./IsFile_test.go"))
	fmt.Println(util.IsFile("./non-existing-file.txt"))
	fmt.Println(util.IsFile("../util")) // Existing but is a dir

	// Output:
	// true
	// false
	// false
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestIsFile(t *testing.T) {
	// Data Provider
	tests := []struct {
		msg    string
		path   string
		expect bool
	}{
		{
			path:   "./IsFile_test.go",
			expect: true,
			msg:    "existing file should return true",
		},
		{
			path:   "./dummy_foo_bar.go",
			expect: false,
			msg:    "unexisting file should return false",
		},
		{
			path:   "./IsFile_test.go/",
			expect: false,
			msg:    "existing file but treating as a dir should be false",
		},
		{
			path:   "../util",
			expect: false,
			msg:    "existing dir but not a file should be false",
		},
		{
			path:   "../util/",
			expect: false,
			msg:    "existing dir but not a file should be false",
		},
		{
			path:   "/bin/bon/ban",
			expect: false,
			msg:    "unexisting file should be false",
		},
		{
			path:   "/path/to/whatever",
			expect: false,
			msg:    "totally unexisting path shoul be false",
		},
	}

	for _, test := range tests {
		actual := util.IsFile(test.path)
		expect := test.expect

		assert.Equal(t, expect, actual, "%s (Target path: %s)", test.msg, test.path)
	}
}
