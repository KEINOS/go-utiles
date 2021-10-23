package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleGenMask() {
	fmt.Printf("%b\n", util.GenMask(0))
	fmt.Printf("%b\n", util.GenMask(1))
	fmt.Printf("%b\n", util.GenMask(2))
	fmt.Printf("%b\n", util.GenMask(4))
	fmt.Printf("%b\n", util.GenMask(11))
	fmt.Printf("%b\n", util.GenMask(64))
	fmt.Printf("%b\n", util.GenMask(65))   // Max is 64
	fmt.Printf("%b\n", util.GenMask(1024)) // Max is 64

	// Output:
	// 0
	// 1
	// 11
	// 1111
	// 11111111111
	// 1111111111111111111111111111111111111111111111111111111111111111
	// 1111111111111111111111111111111111111111111111111111111111111111
	// 1111111111111111111111111111111111111111111111111111111111111111
}

func TestGenMask(t *testing.T) {
	for i := 1; i < 65; i++ {
		mask := util.GenMask(i)
		lenMask := len(fmt.Sprintf("%b", mask))
		t.Logf("%v -> Mask: %v (0b%b)", i, mask, mask)

		require.Equal(t, i, lenMask,
			"the length should be equal %v", mask,
		)
	}

	// Test more than 64
	mask := util.GenMask(65)
	lenMask := len(fmt.Sprintf("%b", mask))

	assert.Equal(t, 64, lenMask, "the max len should be 64")
}
