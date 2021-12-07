package util_test

import (
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/require"
)

func TestParseVersion(t *testing.T) {
	for _, test := range []struct {
		expect map[string]string
		input  string
	}{
		// Go style versioning
		{
			map[string]string{"major": "1", "minor": "2", "patch": "3", "prerelease": "alpha", "build": "12345"},
			"v1.2.3-alpha+12345",
		},
		// Semantic versioning
		{
			map[string]string{"major": "1", "minor": "2", "patch": "3", "prerelease": "alpha", "build": "12345"},
			"1.2.3-alpha+12345",
		},
		// Git tag style versioning
		{
			map[string]string{"major": "1", "minor": "2", "patch": "3", "prerelease": "4", "build": "56789"},
			"v1.2.3-4-56789",
		},
		// Go module style versioning
		{
			map[string]string{"major": "1", "minor": "2", "patch": "3", "prerelease": "", "build": "foobar"},
			"v1.2.3 (foobar)",
		},
		// Variation
		{
			map[string]string{"major": "1", "minor": "0", "patch": "0", "prerelease": "", "build": ""},
			"1",
		},
		{
			map[string]string{"major": "1", "minor": "0", "patch": "0", "prerelease": "", "build": ""},
			"v1",
		},
		{
			map[string]string{"major": "1", "minor": "2", "patch": "0", "prerelease": "", "build": ""},
			"1.2",
		},
		{
			map[string]string{"major": "1", "minor": "2", "patch": "3", "prerelease": "", "build": ""},
			"1.2.3",
		},
		{
			map[string]string{"major": "1", "minor": "2", "patch": "3", "prerelease": "", "build": "12345"},
			"v1.2.3+12345",
		},
		// Issue #10
		{
			map[string]string{"major": "1", "minor": "2", "patch": "3", "prerelease": "", "build": "12345"},
			"v.1.2.3+12345",
		},
	} {
		actual, err := util.ParseVersion(test.input)
		require.NoError(t, err)
		require.Equal(t, test.expect, actual, "input: %s", test.input)
	}
}

func TestParseVersion_invalid_version(t *testing.T) {
	input := "version foo bar 1.2.3"
	_, err := util.ParseVersion(input)

	require.Error(t, err, "non-semantic versioned string should return an error.\nInput: %v", input)
}
