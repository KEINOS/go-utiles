package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleUniqSliceString() {
	data := []string{
		"one",
		"one",
		"two",
		"two",
		"three",
		"three",
	}
	fmt.Println(util.UniqSliceString(data))
	// Output: [one two three]
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestUniqSliceString(t *testing.T) {
	data := []string{
		"one",
		"one",
		"two",
		"two",
		"three",
		"three",
	}
	expect := []string{
		"one",
		"two",
		"three",
	}
	actual := util.UniqSliceString(data)

	assert.Equal(t, expect, actual,
		"the values should be unique and not containing a duplicate value nor changing the order")
}

func TestUniqSliceString_same_order(t *testing.T) {
	data := []string{
		"one",
		"one",
		"two",
		"two",
		"three",
		"three",
	}
	expect := []string{
		"three",
		"one",
		"two",
	}
	actual := util.UniqSliceString(data)

	assert.NotEqual(t, expect, actual, "the order of the values should not change")
}
