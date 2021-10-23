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

func ExampleUIntToBase58() {
	// In base58(BTC), zero becomes "1". See EncodeBase58().
	inputZero := uint(0)
	if encZero, err := util.UIntToBase58(inputZero); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(encZero) // note that the result is in 2 digit -> "11"
	}

	inputTen := uint(10)
	if encTen, err := util.UIntToBase58(inputTen); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(encTen)
	}

	inputHuge := uint(123456789)
	if encHuge, err := util.UIntToBase58(inputHuge); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(encHuge)
	}

	// Output:
	// 11
	// 1B
	// BukQL
}

// This function is used when you need a more accurate checksum in 2 digit string.
func ExampleUIntToBase58_more_accurate() {
	input := uint(3363) // 3363 is the max value of 2 digit Base58 "zz"

	// Encode the checksum to Base58 as a string
	enc, err := util.UIntToBase58(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Encoded checksum: %v (%T)", enc, enc)

	// Output:
	// Encoded checksum: zz (string)
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestUIntToBase58(t *testing.T) {
	input := uint(0)

	enc, err := util.UIntToBase58(input)
	require.NoError(t, err)

	expect := "11" // 11 -> 0x00 -> 0
	actual := enc

	assert.Equal(t, expect, actual,
		"in base58(BTC), zero value should be encoded to '1' and return in two digits minimum")
}

func TestUIntToBase58_mal_multibase_type(t *testing.T) {
	// backup and defer recover
	oldMultibaseBase58BTC := util.MultibaseBase58BTC
	defer func() {
		util.MultibaseBase58BTC = oldMultibaseBase58BTC
	}()

	// Mock MultibaseBase58BTC type. 300 is not defined as multibase enc type
	util.MultibaseBase58BTC = 300

	_, err := util.UIntToBase58(uint(1))
	require.Error(t, err, "un-supported enum (base encode type) should be an error")
}
