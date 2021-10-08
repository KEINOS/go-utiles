package util

// UniqSliceString removes duplicate values of a given slice and returns a slice
// with unique values. The order remains the same as the original.
//
// Issue: https://qiitadon.com/web/statuses/106158855888548864
// Ref: https://qiitadon.com/web/statuses/106158948168528024
func UniqSliceString(input []string) []string {
	flipped := make(map[string]bool)
	uniq := []string{}

	for _, element := range input {
		// Skip duplicate
		if flipped[element] {
			continue
		}

		flipped[element] = true // Map the value of input as key

		uniq = append(uniq, element) // Add if not a dup
	}

	return uniq
}
