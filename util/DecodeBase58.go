package util

import "github.com/multiformats/go-multibase"

// DecodeBase58 takes a encoded string of EncodeBase58 and decodes into a bytes
// buffer.
func DecodeBase58(data string) ([]byte, error) {
	// "z" represents Base58BTC
	_, decoded, err := multibase.Decode("z" + data)

	return decoded, err
}
