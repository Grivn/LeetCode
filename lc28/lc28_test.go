package lc28

import "testing"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i:=0; i<len(haystack); i++ {
		if len(haystack)-i < len(needle) {
			return -1
		}

		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

func TestLC28(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	println(strStr(haystack, needle))
}
