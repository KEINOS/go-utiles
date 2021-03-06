package util

import (
	"github.com/pkg/errors"
	"lukechampine.com/blake3"
)

// HashBLAKE3 returns the hashed value of "input" with length of "lenHash".
// The lenHash must be in the range between 1-1024.
//
// The hash algorithm is based on BLAKE3 so it is fast but NOT suitable for
// cryptographic purposes. Only suitable for hashing a small range of values
// such as IDs or temporary values.
//
// The input will be hashed with BLAKE3 algorithm then encodes it to Base58
// (Base58BTC) and returns the first "lenHash" bytes of the results.
func HashBLAKE3(input string, lenHash int) (hashed string, err error) {
	if lenHash > 1024 || lenHash < 1 {
		return "", errors.Errorf("length error. It should be between 1-1024. Len:%v", lenHash)
	}

	// BLAKE3 hasher
	hasher := blake3.New(1024, nil)

	_, err = hasher.Write([]byte(input))

	if err == nil {
		byteHash := hasher.Sum(nil)

		hashed, err = EncodeBase58(byteHash)
	}

	if err == nil {
		hashed = hashed[0:lenHash]
	}

	return hashed, err
}
