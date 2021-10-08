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
//  Example usage
// ----------------------------------------------------------------------------

func ExampleHashBLAKE3() {
	input := "foo"
	lenHash := 16

	hashed, err := util.HashBLAKE3(input, lenHash)
	if err != nil {
		// Do something with the error
		log.Fatalf("failed to hash: %v", err)
	}

	fmt.Println("Hashed value:", hashed)
	fmt.Println("Length:", len(hashed))
	// Output:
	// Hashed value: 7STCqaLBnDB6EKXi
	// Length: 16
}

// ----------------------------------------------------------------------------
/// Tests
// ----------------------------------------------------------------------------

func TestHashBLAKE3(t *testing.T) {
	for _, testCase := range []struct {
		input  string
		expect string
		len    int
	}{
		{"foo", "7STCq", 5},
		{"foo", "7STCqaLBnD", 10},
		{"foo", "7STCqaLBnDB6EKXi9pKq6WGYp8SKYhQ5tzqT4UNSgj2UT5N9nR", 50},
		{"bar", "6XeEX", 5},
		{"bar", "6XeEX8QwymKZDyrhpovGCWDmLxW4DexDFq9BmhFdknbjVtQ4zw", 50},
	} {
		hashed, err := util.HashBLAKE3(testCase.input, testCase.len)
		require.NoError(t, err)

		assert.Equal(t, testCase.expect, hashed,
			"input: %v\nlen: %v", testCase.input, testCase.len)
	}
}

func TestHashBLAKE3_errors(t *testing.T) {
	_, err := util.HashBLAKE3("foo", 0)
	require.Error(t, err, "the length range should be 1-1024")

	_, err = util.HashBLAKE3("foo", 1025)
	require.Error(t, err, "the length range should be 1-1024")
}
