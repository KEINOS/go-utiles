package util_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-utiles/util"
)

func ExampleDecodeBase58() {
	input := "abcdefg"

	encoded, err := util.EncodeBase58([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(encoded)

	decoded, err := util.DecodeBase58(encoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decoded))

	// Output:
	// 4h3c6xC6Mc
	// abcdefg
}
