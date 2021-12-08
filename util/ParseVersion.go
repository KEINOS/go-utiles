package util

import (
	"errors"
	"strings"

	"golang.org/x/mod/semver"
)

// ParseVersion parses a version string into a mapped data.
//
// It is similar to go's semver package, but it includes a build string as well.
// This function is compatible with git-tagged versions.
//
//     ParseVersion("v1.2.3-alpha-abcde123")
//     // => map[string]string{
//     //      "major": "1",
//     //      "minor": "2",
//     //      "patch": "3",
//     //      "prerelease": "alpha",
//     //      "build": "abcde123",
//     //    }, nil
func ParseVersion(version string) (parsed map[string]string, err error) {
	// Pre-convert v1.2.3 (12345) --> v1.2.3+12345
	result := strings.ReplaceAll(version, " ", "")
	result = strings.ReplaceAll(result, "(", "+")
	result = strings.ReplaceAll(result, ")", "")

	// Trim the leading v, dot and spaces v.1.2.3 --> 1.2.3 (Issue #10 fix)
	result = strings.TrimLeft(result, "v. ")

	// Re-add 'v' prefix for Go semver parse compatibility
	if result[0] != 'v' {
		result = "v" + result
	}

	// Pre-convert git tag's format to Go semver style.
	// v1.2.3-4-5678 --> v1.2.3-4+5678
	if preCount := strings.Count(result, "-"); preCount > 1 {
		result = strings.ReplaceAll(result, "-", "+")
		result = strings.Replace(result, "+", "-", preCount-1)
	}

	// Get build before canonicalize
	build := strings.ReplaceAll(semver.Build(result), "+", "")

	// Canonicalize (add missing minor and patch versions)
	if result = semver.Canonical(result); result == "" {
		return nil, errors.New("invalid version string")
	}

	// Get pre-release
	prerelease := strings.TrimLeft(semver.Prerelease(result), "-")

	// Trim 'v' prefix and strip for minor and patch versions
	result = strings.TrimLeft(result, "v. ")
	splited := strings.Split(strings.ReplaceAll(result, "-", "."), ".")

	return map[string]string{
		"major":      splited[0],
		"minor":      splited[1],
		"patch":      splited[2],
		"prerelease": prerelease,
		"build":      build,
	}, nil
}
