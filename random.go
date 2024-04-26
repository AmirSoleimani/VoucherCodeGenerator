package vcgen

import (
	"math/rand"
	"time"
)

// Initializes the random source
func init() {
	rand.NewSource(time.Now().UnixNano())
}

// return random int in the range min...max
func randomInt(min, max int) int {
	return min + rand.Intn(1+max-min)
}

// return random char string from charset
func randomChar(cs []byte) string {
	return string(cs[randomInt(0, len(cs)-1)])
}
