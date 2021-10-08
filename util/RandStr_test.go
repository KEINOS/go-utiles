package util_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleRandStr() {
	length := 16

	for i := 0; i < 1000; i++ {
		h1 := util.RandStr(length)
		h2 := util.RandStr(length)

		if h1 == h2 {
			log.Fatalf("the result did collide\nh1: %v\nh2: %v\n", h1, h2)
		}
	}

	fmt.Println("ok")
	// Output: ok
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestRandStr_out_of_range(t *testing.T) {
	oldOsExit := util.OsExit
	defer func() {
		util.OsExit = oldOsExit
	}()

	capStatus := 0 // should turn to 1

	util.OsExit = func(code int) {
		capStatus = code
	}

	out := capturer.CaptureStderr(func() {
		_ = util.RandStr(1025)
	})

	assert.Contains(t, out, "length error", "the stderr should contain the reason")
	assert.Equal(t, 1, capStatus, "on error the exit status code should be 1")
}
