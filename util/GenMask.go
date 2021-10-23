package util

// GenMask returns a lenBit length value filled with bit 1.
//
// The lenBit should be between 0-64. Any greater number than 64 will be 64.
//
//     i := util.GenMask(0) // -> 0b0
//     i := util.GenMask(1) // -> 0b1
//     i := util.GenMask(4) // -> 0b1111
//     i := util.GenMask(8) // -> 0b11111111
func GenMask(lenBit int) uint {
	if lenBit > 64 {
		lenBit = 64
	}

	result := uint(0)

	for ii := 0; ii < lenBit; ii++ {
		result = (result << 1) + 1
	}

	return result
}
