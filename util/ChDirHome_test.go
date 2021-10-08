package util_test

import (
	"fmt"
	"log"
	"os"

	"github.com/KEINOS/go-utiles/util"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleChDirHome() {
	// Move to current user home dir and defer moving back to original
	funcReturn := util.ChDirHome()
	defer funcReturn()

	/* Get dir infos to check if the current dir is user's home */
	pathDirHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	pathDirCurr, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Assert
	if pathDirCurr == pathDirHome {
		fmt.Println("moved to user's home dir")
	} else {
		log.Fatalf("failed to move dir. Home: %v, Current: %v", pathDirHome, pathDirCurr)
	}

	// Output: moved to user's home dir
}
