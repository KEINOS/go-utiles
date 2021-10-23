package util_test

import (
	"fmt"

	"github.com/KEINOS/go-utiles/util"
)

func ExampleConvBytes2Uint() {
	// Bytes in big endian
	input := []byte{0, 0, 0, 0, 0, 0, 48, 57} // = uint64(12345)

	resultUint64 := util.ConvBytes2Uint(input)
	fmt.Println(resultUint64)

	// To convert uint64 to []byte use ConvUint2Bytes().
	resultByteSlice := util.ConvUint2Bytes(resultUint64)
	fmt.Println(resultByteSlice)

	// Output:
	// 12345
	// [48 57]
}

func ExampleConvUint2Bytes() {
	input := uint(12345) // 0x30 0x39

	resultByteSlice := util.ConvUint2Bytes(input)

	fmt.Println(resultByteSlice)
	fmt.Printf("%#v\n", resultByteSlice)

	resultUint64 := util.ConvBytes2Uint(resultByteSlice)
	fmt.Println(resultUint64)

	// Output:
	// [48 57]
	// []byte{0x30, 0x39}
	// 12345
}

func ExampleConvUint2Bytes_negative_value() {
	input := -123456789

	resultByteSlice := util.ConvUint2Bytes(uint(input)) // note the uint conversion
	resultUint64 := util.ConvBytes2Uint(resultByteSlice)

	fmt.Println(input)
	fmt.Println(int64(resultUint64)) // note the int64
	// Output:
	// -123456789
	// -123456789
}
