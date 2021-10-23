package util_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Examples
// ----------------------------------------------------------------------------

func ExampleSUM() {
	input := "foo bar"
	sum8 := uint(0b11111111) // 8bit mask = 255 = checksum between 1-255

	checksum := util.SUM(sum8, input)
	fmt.Printf("%d (0x%x, %T)\n", checksum, checksum, checksum)

	// Output:
	// 156 (0x9c, uint)
}

func ExampleSUM_more_accurate() {
	const (
		input     = "foo bar"
		sumBase58 = uint(3363) // 3363 is the max number of 2 digit Base58 = "zz"
	)

	// Create checksum
	checksum := util.SUM(sumBase58, input)
	fmt.Printf("Checksum: %v (0x%x, 0b%b, %T)\n", checksum, checksum, checksum, checksum)

	// Encode to Base58
	enc, err := util.UIntToBase58(checksum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Base58 encoded: %v (%T)\n", enc, enc)
	// Output:
	// Checksum: 666 (0x29a, 0b1010011010, uint)
	// Base58 encoded: CV (string)
}

func ExampleSUM_with_verify() {
	input := util.RandStr(1024) // 1024 char length random string

	sum8 := uint(255) // checksum for max 8bit = 0b11111111 = 0d255

	checksum := util.SUM(sum8, input)

	// Verify
	if util.VerifySUM(sum8, input, checksum) {
		fmt.Print("verify success! checksum of the input is valid")
	}

	// Output: verify success! checksum of the input is valid
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestSUM_max_checksum(t *testing.T) {
	const (
		sum8  = uint(0b11111111) // 255
		input = "RYrSdJfNcrsWRZk3XKUcHo86orGLX4VDTPnGoxE53caMU3bMMd" +
			"ret6syH5xeFFm98tLnBQntnQGnPyCuj2J8hyvm5QwQ6Bj88K6V"
	)

	expect := uint(255)
	actual := util.SUM(sum8, input)

	assert.Equal(t, expect, actual)
}

func TestSUM_min_checksum(t *testing.T) {
	const (
		sum8  = uint(0b11111111) // 255
		input = "3oaBzKVM1eWdzDKHSyBxEEzBRmDVMXVVdrHCqhWqq1owPd1V9y" +
			"2sfgPAqupjZB2gWBo6PFm2Ee3dqcXV9dtP74TanvqoFgLdtwmy"
	)

	expect := uint(1)
	actual := util.SUM(sum8, input)

	assert.Equal(t, expect, actual)
}

func TestSUM_not_greater_than_the_mask(t *testing.T) {
	max := 3363 // 3363 is the max number of 2 digit Base58 = "zz"

	for i := 1; i <= max; i++ {
		input := util.RandStr(100) // random string of 100 chars
		lenBit := uint(i)          // mask

		checksum := util.SUM(lenBit, input)

		require.GreaterOrEqual(t, lenBit, checksum,
			"the checksum should not be greater than the mask (lenBit)")
	}
}
