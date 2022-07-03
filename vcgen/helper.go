package vcgen

import (
	"math"
	"math/rand"
	"strings"
)

//return random int in the range min...max
func randomInt(min, max int) int {
	return min + rand.Intn(1+max-min)
}

//return random char string from charset
func randomChar(cs []byte) string {
	return string(cs[randomInt(0, len(cs)-1)])
}

//repeat string with one str (#)
func repeatStr(count uint16, str string) string {
	return strings.Repeat(str, int(count))
}

func numberOfChar(str, char string) uint16 {
	return uint16(strings.Count(str, char))
}

func isFeasible(charset, pattern, char string, count uint16) bool {
	ls := numberOfChar(pattern, char)
	return math.Pow(float64(len(charset)), float64(ls)) >= float64(count)
}
