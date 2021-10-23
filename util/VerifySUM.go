package util

import "fmt"

func VerifySUM(mask uint, input string, sum uint) bool {
	filter := GenMask(len(fmt.Sprintf("%b", mask))) // create N digit filter

	sumComp := summation(input)
	sumComp = filter - (sumComp % mask)
	sumComp += sum

	filtered := sumComp & filter // Use the last N bit(filter) of the sum

	return filtered == 0
}
