package vcgen

import (
	"math"
	"strings"
)

// repeat string with one str (#)
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
