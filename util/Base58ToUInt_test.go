package util_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-utiles/util"
)

func ExampleBase58ToUInt() {
	const (
		encValue = "zz"       // Base58 encoded of value 3363
		expect   = uint(3363) // unsigned expect value
	)

	// Decode to uint
	actual, err := util.Base58ToUInt(encValue)
	if err != nil {
		log.Fatal(err)
	}

	if expect == actual {
		fmt.Println("it is a valid checksum!")
	}

	// Output: it is a valid checksum!
}
