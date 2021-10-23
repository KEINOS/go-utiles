package util

import "encoding/binary"

// ConvUint2Bytes converts uint to []byte (big endian).
//
// To conver []byte to uint use ConvBytes2Uint().
func ConvUint2Bytes(i uint) []byte {
	byteBox := make([]byte, 8) // 8x8=64

	binary.BigEndian.PutUint64(byteBox, uint64(i))

	foundNum := false
	result := []byte{}

	for _, data := range byteBox {
		if data == 0 && !foundNum {
			continue
		}

		foundNum = true

		result = append(result, data)
	}

	return result
}
