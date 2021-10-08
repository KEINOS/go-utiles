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

func ExampleFmtStructPretty_json() {
	data := struct {
		Foo string `json:"foo" mapstructure:"foo"`
	}{
		Foo: "bar",
	}

	prettyJSON := util.FmtStructPretty(data)

	fmt.Println(prettyJSON)
	// Output: {
	//   "foo": "bar"
	// }
}

func ExampleFmtStructPretty_slice() {
	data := []string{
		"foo",
		"bar",
	}

	prettyJSON := util.FmtStructPretty(data)

	fmt.Println(prettyJSON)
	// Output: [
	//   "foo",
	//   "bar"
	// ]
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestFmtStructPretty(t *testing.T) {
	data := struct {
		Foo string
	}{
		Foo: "bar",
	}

	{
		expect := "{\n*****  \"Foo\": \"bar\"\n*****}"
		actual := util.FmtStructPretty(data, "*****")

		assert.Equal(t, expect, actual)
	}
	{
		expect := "{\n    \"Foo\": \"bar\"\n  }"
		actual := util.FmtStructPretty(data, "  ")

		assert.Equal(t, expect, actual)
	}
}

func TestFmtStructPretty_tagged_struct(t *testing.T) {
	data := struct {
		Foo string `json:"foo" mapstructure:"foo"`
	}{
		Foo: "bar",
	}

	{
		expect := "{\n*****  \"foo\": \"bar\"\n*****}"
		actual := util.FmtStructPretty(data, "*****")

		assert.Equal(t, expect, actual)
	}
	{
		expect := "{\n    \"foo\": \"bar\"\n  }"
		actual := util.FmtStructPretty(data, "  ")

		assert.Equal(t, expect, actual)
	}
}
