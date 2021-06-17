package lc5

import (
	"testing"
)

func longestPalindrome(s string) string {
	p := ""
	for subLen:=len(s); subLen>0; subLen-- {
		subSet := getSubStr(s, subLen)

		p = checkP(subSet)

		if len(p) > 0 {
			break
		}
	}

	return p
}

func getSubStr(str string, subLen int) []string {
	var subSet []string

	for i:=0; i<len(str); i++ {
		if len(str)-i < subLen {
			break
		}

		sub := str[i:i+subLen]

		subSet = append(subSet, sub)
	}

	return subSet
}

func checkP(subSet []string) string {
	for _, sub := range subSet {
		good := true
		for i:=0; i<len(sub); i++ {
			j := len(sub)-i-1

			if sub[i] != sub[j] {
				good = false
				break
			}

			if i == j || i+1 == j {
				break
			}
		}

		if good {
			return sub
		}
	}

	return ""
}

func TestLC5(t *testing.T) {
	str := "cbbd"

	p := longestPalindrome(str)

	println(p)
}
