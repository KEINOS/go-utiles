package util

import "fmt"

// SUM8 returns the sum8 algorithm checksum of input as a 1 Byte(8 bit, 2 chars)
// hex string.
//
// This value can be verified by VeifySUM8 function by attaching the checksum
// value to the input as below.
//
//     input := "foo bar buz"
//     checksum := util.SUM8(input)
//     if util.VerifySUM8(input + checksum) {
//         fmt.Println("ok")
//     }
//
//     https://play.golang.org/p/HPjGBJt7f_6
func SUM8(input string) string {
	// Create bit mask for the last 8 bits(11...1100000000)
	const mask = (^uint(0) >> 8) << 8

	sum := uint(0)

	getLast8bit := func(data uint) uint8 {
		return uint8(data &^ mask)
	}

	for _, B := range input {
		sum += uint(B)

		if sum >= 512 {
			sum = uint(getLast8bit(sum))
		}
	}

	last8bit := getLast8bit(sum)
	flipped := uint(^last8bit) + 1
	checksum := getLast8bit(flipped)

	return fmt.Sprintf("%02x", checksum)
}
