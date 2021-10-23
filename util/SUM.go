package util

import "fmt"

// SUM returns the checksum of the input based on 2's complement of the sum with
// max length of the mask.
//
// The returned sum will be between 1 - mask. For example if the mask is 255, then
// the checksum will be between 1-255.
//
// To verify the checksum with the input, use VerifySUM() function.
func SUM(mask uint, input string) uint {
	filter := GenMask(len(fmt.Sprintf("%b", mask))) // create N digit filter

	sum := summation(input)
	sum = filter - (sum % mask)

	checksum := ^sum + 1

	checksum = checksum & filter // Use the last N bit(filter) of the sum

	return checksum
}

// It returns the summation of every 8 bit data of the input string.
func summation(input string) uint {
	sum := uint(0)

	// sum up every 8bit chunk data
	for _, B := range []byte(input) {
		sum += uint(B)
	}

	return sum
}
