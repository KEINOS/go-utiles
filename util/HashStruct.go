package util

// HashStruct returns the hash value of the input struct with the given length.
//
// Note that the hash value is only for change detection purposes and NOT to detect falsification.
func HashStruct(input interface{}, lenHash int) (string, error) {
	inputJSON := FmtStructPretty(input)

	return HashBLAKE3(inputJSON, lenHash)
}
