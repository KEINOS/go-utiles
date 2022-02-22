package util

import (
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getModName_issue_1(t *testing.T) {
	for _, test := range []struct {
		in  string
		out string
	}{
		{"github.com/jedib0t/go-pretty/v6", "go-pretty/v6"},
		{"github.com/klauspost/cpuid/v2", "cpuid/v2"},
		{"lukechampine.com/blake3", "blake3"},
		{"gopkg.in/yaml.v3", "yaml.v3"},
	} {
		m := &debug.Module{
			Path:    test.in,
			Version: "0.0.0",    // dummy
			Sum:     "ffffffff", // dummy
		}

		expect := test.out
		actual := getModName(m)

		assert.Equal(t, expect, actual)
	}
}
