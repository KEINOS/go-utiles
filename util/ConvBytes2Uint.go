package util

import "encoding/binary"

// ConvBytes2Uint converts []byte (big endian) to uint.
//
// To convert uint to []byte use ConvUint2Bytes().
func ConvBytes2Uint(input []byte) uint {
	lenInput := len(input)

	if lenInput == 8 {
		return uint(binary.BigEndian.Uint64(input))
	}

	byteBox := make([]byte, 8) // 8x8=64

	for i, data := range input {
		n := 8 - lenInput + i

		byteBox[n] = data
	}

	return uint(binary.BigEndian.Uint64(byteBox))
}
