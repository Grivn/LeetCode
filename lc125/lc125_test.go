package lc125

import (
	"fmt"
	"testing"
)

func isPalindrome(s string) bool {
	set := strToSet(s)

	for i := range set {
		if set[i] == set[len(set)-1-i] {
			continue
		} else {
			return false
		}
	}

	return true
}

func strToSet(s string) []int {
	var ret []int

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			ret = append(ret, int(c))
			continue
		}

		if c >= 'A' && c <= 'Z' {
			ret = append(ret, int(c+32))
			continue
		}

		if c >= '0' && c <= '9' {
			ret = append(ret, int(c))
			continue
		}
	}

	return ret
}

func TestLC125(t *testing.T) {
	str := "az,AZ"

	fmt.Println(isPalindrome(str))
}
