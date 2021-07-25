package lc91

import (
	"fmt"
	"testing"
)

// NOTE!!!
// important DP problem

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	result := make([]int, len(s))
	result[0] = 1

	for i, c := range s {
		if i == 0 {
			if c == '0' {
				return 0
			}
			result[i] = 1
			continue
		}

		if i == 1 {
			if c == '0' {
				result[i] = 0
			} else {
				result[i] = result[i-1]
			}

			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
				result[i] += 1
			}
			continue
		}

		if c == '0' {
			result[i] = 0
		} else {
			result[i] = result[i-1]
		}

		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			result[i] += result[i-2]
		}
	}

	fmt.Println(result)
	return result[len(s)-1]
}

func TestLC91(t *testing.T) {
	str := "226"
	println(numDecodings(str))
}
