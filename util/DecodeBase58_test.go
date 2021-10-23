package util_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleDecodeBase58() {
	input := "abcdefg"

	// Encode
	encoded, err := util.EncodeBase58([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encoded:", encoded)

	// Decode
	decoded, err := util.DecodeBase58(encoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decoded:", string(decoded))

	// Output:
	// Encoded: 4h3c6xC6Mc
	// Decoded: abcdefg
}

func TestDecodeBase58(t *testing.T) {
	for _, testData := range []struct {
		input  string
		expect uint
	}{
		{"z", 57},  // 57 is the max value for 1 digit string(since it's base 58)
		{"1z", 57}, // 1 is zero in base 58
		{"1", 0},
		{"11", 0},
		{"zy", 3362},
		{"zz", 3363}, // 3363 is the max value for 2 digit string
		{"211", 3364},
	} {
		actualByte, err := util.DecodeBase58(testData.input)
		require.NoError(t, err)

		expect := testData.expect
		actual := util.ConvBytes2Uint(actualByte)

		assert.Equal(t, expect, actual)
	}
}

func TestDecodeBase58_malencoding(t *testing.T) {
	for _, input := range "0OIl" {
		_, err := util.Base58ToUInt(string(input))
		require.Error(t, err, "Base58 should not include '0', 'O', 'I', 'l'")
	}
}
