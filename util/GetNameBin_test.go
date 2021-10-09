package util_test

import (
	"fmt"

	"github.com/KEINOS/go-utiles/util"
)

func ExampleGetNameBin() {
	nameBin := util.GetNameBin()

	fmt.Println(nameBin)

	// Output: util.test
}
