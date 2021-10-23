package util

import "fmt"

// VerifySUM returns true if the sum is a valid checksum of the input with the given mask.
// The sum value should be created via SUM() function with the same mask value.
func VerifySUM(mask uint, input string, sum uint) bool {
	filter := GenMask(len(fmt.Sprintf("%b", mask))) // create N digit filter

	sumComp := summation(input)
	sumComp = filter - (sumComp % mask)
	sumComp += sum

	filtered := sumComp & filter // Use the last N bit(filter) of the sum

	return filtered == 0
}
