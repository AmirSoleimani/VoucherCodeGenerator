package vcgen

import (
	"strings"
	"testing"
)

func TestRandomInt(t *testing.T) {
	tables := []struct {
		min, max int
	}{
		{1, 10},
		{100, 200},
	}

	for _, v := range tables {
		n := randomInt(v.min, v.max)
		if n > v.max || n < v.min {
			t.Fail()
		}
	}
}

func TestRandomChar(t *testing.T) {
	tables := []struct {
		chars []byte
	}{
		{[]byte("abcljasdlkjasd")},
		{[]byte("kqweqwrkjhasfn")},
	}

	for _, v := range tables {
		c := randomChar(v.chars)
		if !strings.Contains(string(v.chars), c) {
			t.Fail()
		}
	}
}
