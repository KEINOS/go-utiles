package util

import (
	"fmt"
	"math/rand"
	"time"
)

// RandStr returns a random unique string with the given length.
// The length range must be between 1-1024. Otherwise it will os.Exit with status 1.
//
// Note that, it is a pseudo-random string generator and unsuitable for security-sensitive work.
func RandStr(length int) string {
	//nolint:gosec // not for cryptographically secure random purposes
	salty := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := fmt.Sprintf("%v:%v", time.Now().UnixNano(), salty.Float64())

	h, err := HashBLAKE3(str, length)
	ExitOnErr(err)

	return h
}
