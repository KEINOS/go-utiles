package util

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

// FmtStructPretty formats JSON string or an object into pretty-indented
// JSON-strings.
//
// If a prefix is provided then it will add the prefix to each line.
func FmtStructPretty(val interface{}, prefixes ...string) string {
	var (
		prefix      = ""
		paddingLeft = 2
		indent      = strings.Repeat(" ", paddingLeft)
	)

	if len(prefixes) > 0 {
		prefix = strings.Join(prefixes, "")
	}

	transformer := text.NewJSONTransformer(prefix, indent)

	return transformer(val)
}
