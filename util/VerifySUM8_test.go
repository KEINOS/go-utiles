package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
)

func ExampleVerifySUM8() {
	input := "abcdefghijk" // target data
	checksum := "9e"       // checksum value (result of: util.SUM8(input))

	// Value with sum8 checksum
	data := input + checksum

	// Verify
	if util.VerifySUM8(data) {
		fmt.Println("checksum is valid")
	}

	// Output: checksum is valid
}

func TestVerifySUM8_not_hex(t *testing.T) {
	input := "aaaa"
	checksum := "zz"

	assert.False(t, util.VerifySUM8(input+checksum), "invalid checksum should return an error")
}

func TestVerifySUM8_too_short(t *testing.T) {
	for _, data := range []struct {
		input  string
		expect bool
	}{
		{"a", false},    // too short
		{"aa", false},   // too short
		{"aaa", false},  // invalid checksum
		{"a9f", true},   // "9f" is the checksum of "a"
		{"aaaa", false}, // invalid checksum
		{"aa3e", true},  // "3e" is the checksum of "aa"
	} {
		input := data.input

		expect := data.expect
		actual := util.VerifySUM8(input)

		assert.Equal(t, expect, actual,
			"less than 3 chars should return an error.\nInput: %v", input)
	}
}
