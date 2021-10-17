package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/require"
)

func ExampleSUM8() {
	// Get sum8 checksum
	fmt.Println(util.SUM8("abcdefghijk"))

	// Output: 9e
}

func TestSUM8(t *testing.T) {
	for i := 1; i < 1024; i++ {
		input := util.RandStr(i)
		checksum := util.SUM8(input)

		require.True(t, util.VerifySUM8(input+checksum),
			"checksum not match. Input: %v, Checksum: %v, Len: %v", input, checksum, i)
	}
}

func TestSUM8_fail(t *testing.T) {
	input := util.RandStr(1024)
	checksum := util.SUM8(input)
	dummyInput := util.RandStr(1024)

	require.False(t, util.VerifySUM8(dummyInput+checksum),
		"checksum should not match. Input: %v, Checksum: %v", input, checksum)
}
