package util

import (
	"path/filepath"
	"strings"
)

// IsNameFileJSON returns true if name is a file path and ends with ".json".
func IsNameFileJSON(name string) bool {
	return filepath.Ext(strings.ToLower(name)) == ".json"
}
