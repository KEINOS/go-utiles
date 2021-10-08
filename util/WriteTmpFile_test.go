package util_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-utiles/util"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleWriteTmpFile() {
	data := "foo bar"

	pathFile, deferCleanUp, err := util.WriteTmpFile(data)
	if err != nil {
		log.Fatal(err)
	}

	defer deferCleanUp()

	read, err := util.ReadFile(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(read))
	// Output: foo bar
}
