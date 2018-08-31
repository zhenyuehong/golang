package main

import "testing"

//表格测试
func TestSubStr(t *testing.T) {
	test := []struct {
		s   string
		ans int
	}{
		//normal case
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},

		//edge case
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},

		//chinese support
		{"说我可以考虑下", 8},
		{"一二三二一", 3},
	}

	for _, tt := range test {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for intput %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}
