package util

import (
	"strings"

	"github.com/MakeNowJust/heredoc"
)

// HereDoc returns an un-indented string such as here-document-like format.
// Useful for help messages to print.
//
// If indents were given it will use as a prefix of each line.
func HereDoc(input string, indents ...string) string {
	input = heredoc.Doc(input)

	if len(indents) == 0 {
		return input
	}

	lines := strings.Split(input, "\n")
	indent := strings.Join(indents, "")

	for i, line := range lines {
		indented := indent + line

		// trim trailing whitespace if line is empty
		if strings.TrimSpace(line) == "" {
			indented = strings.TrimSpace(indented)
		}

		lines[i] = indented
	}

	return strings.Join(lines, "\n")
}
