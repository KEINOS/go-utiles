package util

import (
	"encoding/hex"
)

// VerifySUM8 は inputWithCheckSum の末尾 2 桁に SUM8 のチェックサムが付いた文
// 字列が有効な場合に true を返します。
func VerifySUM8(inputWithCheckSum string) bool {
	// 最後の 8 ビット取得用のマスク（11...1100000000)
	const mask = (^uint(0) >> 8) << 8

	lenData := len(inputWithCheckSum)
	byteData := []byte(inputWithCheckSum)

	if lenData < 3 {
		return false
	}

	// 最後の 2 文字（16 進数の文字列 1 バイトぶん）を数値に変換
	checkSum, err := hex.DecodeString(string(byteData[lenData-2 : lenData]))
	if err != nil {
		return false
	}

	getLast8bit := func(input uint) uint8 {
		return uint8(input &^ mask)
	}

	sum := uint(0)
	for _, B := range byteData[0 : lenData-2] {
		sum += uint(B)

		if sum >= 512 {
			sum = uint(getLast8bit(sum))
		}
	}

	sum += uint(checkSum[0]) // チェックサム（補数）を足す

	return getLast8bit(sum) == 0
}
