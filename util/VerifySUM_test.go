package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Examples
// ----------------------------------------------------------------------------

func ExampleVerifySUM() {
	const (
		input  = "abcdefghijk"    // target data
		bitLen = uint(0b11111111) // 0b11111111 = 255
	)

	// Create checksum between 1-255
	checksum := util.SUM(bitLen, input)
	fmt.Printf("Checksum is: %v (0b%b)\n", checksum, checksum)

	// Verify
	if util.VerifySUM(bitLen, input, checksum) {
		fmt.Println("Check result: ok")
	} else {
		fmt.Println("Check result: ng")
	}

	// Output:
	// Checksum is: 103 (0b1100111)
	// Check result: ok
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestVerifySUM(t *testing.T) {
	for mask := uint(1); mask <= 3363; mask++ {
		input := util.RandStr(100)        // random string with 100 char length
		checksum := util.SUM(mask, input) // create checksum

		// Verify
		require.True(t, util.VerifySUM(mask, input, checksum),
			"checksum miss match.\nInput: %v\nChecksum: %v\nMask len: %v\n", input, checksum, mask)
	}
}

func TestVerifySUM_false(t *testing.T) {
	input := "foo bar"
	mask := uint(0b1111) // sum between 1-15
	sum := util.SUM(mask, input)

	inputDummy := "foo bar buz"

	assert.False(t, util.VerifySUM(mask, inputDummy, sum),
		"verifing the different input should be false")
}
