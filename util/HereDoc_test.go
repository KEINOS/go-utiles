package util_test

import (
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
)

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
