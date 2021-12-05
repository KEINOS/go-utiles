package util

import (
	"strings"

	"github.com/MakeNowJust/heredoc"
)

// HereDoc returns an un-indented string as here-document-like format.
//
// If indents was given, it first will un-indent then indents with the given
// joined indent.
func HereDoc(input string, indents ...string) string {
	input = heredoc.Doc(input)

	if len(indents) == 0 {
		return input
	}

	lines := strings.Split(input, "\n")
	indent := strings.Join(indents, "")

	for i, line := range lines {
		if strings.TrimSpace(line) != "" {
			lines[i] = indent + line
		}
	}

	return strings.Join(lines, "\n")
}
