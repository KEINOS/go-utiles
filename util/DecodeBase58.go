package util

import (
	"github.com/multiformats/go-multibase"
	"github.com/pkg/errors"
)

// DecodeBase58 takes a encoded string of EncodeBase58 and decodes into a bytes
// buffer.
func DecodeBase58(data string) ([]byte, error) {
	// "z" represents Base58BTC
	_, decoded, err := multibase.Decode("z" + data)
	if err != nil {
		err = errors.Wrap(err, "failed to decode base58 encoded data")
	}

	return decoded, err
}
