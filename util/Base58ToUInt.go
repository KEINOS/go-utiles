package util

import "github.com/pkg/errors"

// Base58ToUInt returns the decoded value of enc. The enc value must be Base58
// encoded.
//
// This function is basically used for human readable checksum by encoding/decoding
// the checksum values to Base58 and vice versa.
func Base58ToUInt(enc string) (uint, error) {
	raw, err := DecodeBase58(enc)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert base58 to uint")
	}

	return ConvBytes2Uint(raw), nil
}
