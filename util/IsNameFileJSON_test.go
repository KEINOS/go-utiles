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

func ExampleIsNameFileJSON() {
	if target := "/foo/bar/buz.json"; util.IsNameFileJSON(target) {
		fmt.Println(target, "is a JSON file")
	}

	// Output: /foo/bar/buz.json is a JSON file
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestIsNameFileJSON(t *testing.T) {
	{
		nameFile := "sample.json"
		result := util.IsNameFileJSON(nameFile)

		assert.True(t, result)
	}
	{
		nameFile := "sample.JSON"
		result := util.IsNameFileJSON(nameFile)

		assert.True(t, result)
	}
	{
		nameFile := "./path/to/sample.json"
		result := util.IsNameFileJSON(nameFile)

		assert.True(t, result)
	}
	{
		nameFile := "./path/to/sample"
		result := util.IsNameFileJSON(nameFile)

		assert.False(t, result, "if file name does not end with '.json' then it should return false")
	}
}
