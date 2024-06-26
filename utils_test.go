package vcgen

import "testing"

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
