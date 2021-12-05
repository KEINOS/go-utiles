package util_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
)

func ExampleHereDoc() {
	msg := util.HereDoc(`
        Here Title
            Here description
	`)

	fmt.Println(msg)

	// Output:
	// Here Title
	//     Here description
}

func ExampleHereDoc_optional_indentation() {
	input := `
        Here Title
            Here description
`
	indent := "> " // prefix of each line

	fmt.Println(util.HereDoc(input, indent))

	// Output:
	// > Here Title
	// >     Here description
	// >
}

func TestHereDoc(t *testing.T) {
	input := `
		Level1
			Level2
	`
	expect := "Level1\n\tLevel2\n"
	actual := util.HereDoc(input)
	assert.Equal(t, expect, actual)
}

func TestHereDoc_optional_indent(t *testing.T) {
	input := `
		Level1
			Level2
	`

	expect := "  Level1\n  \tLevel2\n"
	actual := util.HereDoc(input, " ", " ")
	assert.Equal(t, expect, actual)
}
