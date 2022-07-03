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

func TestRepeatStr(t *testing.T) {
	tables := []struct {
		count uint16
		str   string
	}{
		{10, "a"},
		{30, "a"},
		{0, "a"},
	}

	for _, v := range tables {
		rs := repeatStr(v.count, v.str)
		if uint16(len(rs)) != v.count {
			t.Fail()
		}
	}
}

func TestIsFeasible(t *testing.T) {
	tables := []struct {
		charset, pattern string
		count            uint16
		wants            bool
	}{
		{"abcdefghijk", "##-#####-###", 10000, true},
		{"abcdefghijk", "##-#####-###", 10000, true},
		{"abcdefghijk", "##-#", 10000, false},
	}

	for _, v := range tables {
		b := isFeasible(v.charset, v.pattern, "#", v.count)
		if b != v.wants {
			t.Fail()
		}
	}
}
