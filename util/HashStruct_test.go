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

func ExampleHashStruct() {
	data := struct {
		Foo string
		Bar int
	}{
		Foo: "hoge fuga",
		Bar: 1,
	}

	hash1, err := util.HashStruct(data, 16) // 16 char length
	if err != nil {
		log.Fatalf("hash error: %v", err)
	}

	fmt.Println("Hash value before change:", hash1)

	data.Bar = 2 // Change value

	hash2, err := util.HashStruct(data, 16) // 16 char length
	if err != nil {
		log.Fatalf("hash error: %v", err)
	}

	fmt.Println("Hash value after change :", hash2)

	// Output:
	// Hash value before change: 4KcWDdX1qXnGBV4U
	// Hash value after change : 6aESjhTWhk3Tv91h
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestHashStruct_struct_data(t *testing.T) {
	dummyStruct := []struct {
		fieldOne string
		fieldTwo int
	}{
		{"foo", 1},
		{"bar", 2},
	}

	expect := "SAEHJkpjEQfSQRMB"
	actual, err := util.HashStruct(dummyStruct, 16)

	require.NoError(t, err)
	assert.Equal(t, expect, actual)
}

func TestHashStruct_string_data(t *testing.T) {
	dummyStruct := "foo bar"

	expect := "4MiNTebs2jc8yWF6"
	actual, err := util.HashStruct(dummyStruct, 16)

	require.NoError(t, err)
	assert.Equal(t, expect, actual)
}

func TestHashStruct_slice_data(t *testing.T) {
	dummyStruct := []string{
		"foo",
		"bar",
	}

	expect := "3xfQcEGg6n2wBLmF"
	actual, err := util.HashStruct(dummyStruct, 16)

	require.NoError(t, err)
	assert.Equal(t, expect, actual)
}
