package util

import "github.com/multiformats/go-multibase"

// EncodeBase58 returns the Base58 encoded string using Multibase Base58BTC
// format without encode type prefix "z".
func EncodeBase58(input []byte) (string, error) {
	hashed, err := multibase.Encode(multibase.Base58BTC, input)

	if err == nil {
		// Trim the 1st base-encode type character "z"
		hashed = hashed[1:]
	}

	return hashed, err
}
