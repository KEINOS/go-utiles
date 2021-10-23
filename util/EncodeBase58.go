package util

import "github.com/multiformats/go-multibase"

// MultibaseBase58BTC is a copy of multibase.Base58BTC to ease mock
// multibase.Base58BTC for testing.
//
// This library uses MultibaseBase58BTC instead of multibase.Base58BTC, assign
// a dummy function to mock it's behavior.
var MultibaseBase58BTC multibase.Encoding = multibase.Base58BTC

// EncodeBase58 returns the Base58 encoded string using Multibase Base58BTC
// format without the encode type prefix "z".
//
// The used chars are:
//     "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
//     See: https://en.bitcoin.it/wiki/Base58Check_encoding
func EncodeBase58(input []byte) (string, error) {
	hashed, err := multibase.Encode(MultibaseBase58BTC, input)
	if err == nil {
		// Trim the 1st base-encode type character "z"
		hashed = hashed[1:]
	}

	return hashed, err
}
