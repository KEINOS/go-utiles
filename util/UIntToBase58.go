package util

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// UIntToBase58 returns Base58(BTC) encoded string of the given uint value.
// Note that this function returns in 2 digit minimum. Such as 0d0 -> "11".
//
// This function is basically used for human readable checksum by encoding/decoding
// the checksum values to Base58 and vice versa.
func UIntToBase58(value uint) (string, error) {
	result, err := EncodeBase58(ConvUint2Bytes(value))
	if err != nil {
		return "", errors.Wrap(err, "failed to convert uint to base58")
	}

	// Make 1 digit result to two, such as "z" -> "0z" then replace the "0" to
	// "1" since zero in Base58 is "1".
	return strings.ReplaceAll(fmt.Sprintf("%02s", result), "0", "1"), nil
}
