package util_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleEncodeBase58() {
	input := "abcdefg"

	result, err := util.EncodeBase58([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	// Output: 4h3c6xC6Mc
}

func TestEncodeBase58(t *testing.T) {
	// For the pattern see:
	// https://en.bitcoin.it/wiki/Base58Check_encoding
	expectOrder := "123456789" +
		"ABCDEFGHJKLMNPQRSTUVWXYZ" +
		"abcdefghijkmnopqrstuvwxyz"

	assert.Equal(t, 58, len(expectOrder))

	for i, expect := range expectOrder {
		actual, err := util.EncodeBase58([]byte{byte(i)})
		require.NoError(t, err)

		assert.Equal(t, string(expect), actual)
	}

	// Over 58
	val, err := util.EncodeBase58([]byte{byte(58)})

	require.NoError(t, err)
	assert.Equal(t, "21", val)
}

func TestEncodeBase58_3363_is_max(t *testing.T) {
	for _, testData := range []struct {
		expect string
		input  uint
	}{
		{"zy", 3362},
		{"zz", 3363}, // 3363 is the max value for 2 digit string
		{"211", 3364},
	} {
		expect := testData.expect
		actual, err := util.EncodeBase58(util.ConvUint2Bytes(testData.input))
		require.NoError(t, err)

		assert.Equal(t, expect, actual)
	}
}
