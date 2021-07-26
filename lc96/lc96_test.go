package lc96

import (
	"testing"
)

func numTrees(n int) int {

	result := make([]int, n+1)

	result[0] = 1

	result[1] = 1

	for index:=1; index<=n; index++ {

		total := 0

		for ptr:=1; ptr<=index; ptr++ {
			lCount := ptr-1
			rCount := index-ptr
			total += result[lCount]*result[rCount]
		}

		result[index] = total
	}
	return result[n]
}

func TestLC96(t *testing.T) {
	println(numTrees(4))
}
