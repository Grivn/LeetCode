package lc792

import "testing"

func numMatchingSubseq(s string, words []string) int {

	count := 0

	for _, word := range words {
		if check(s, word) {
			count++
		}
	}

	return count
}

func check(s string, word string) bool {
	i := 0
	j := 0
	for {
		if i == len(s) {
			break
		}

		if s[i] == word[j] {
			j++
			if len(word) == j {
				break
			}
		}

		i++
	}

	return j==len(word)
}

func TestLC792(t *testing.T) {
	s := "dsahjpjauf"
	words := []string{"ahjpjau","ja","ahbwzgqnuk","tnmlanowax"}

	println(numMatchingSubseq(s, words))
}
