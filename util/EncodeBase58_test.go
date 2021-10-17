package util_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-utiles/util"
)

func ExampleEncodeBase58() {
	input := "abcdefg"

	result, err := util.EncodeBase58([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	// Output: 4h3c6xC6Mc
}
