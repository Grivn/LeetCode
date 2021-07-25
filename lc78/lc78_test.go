package lc78

import (
	"fmt"
	"testing"
)

// NOTE!!!
// the memory copy for prefix.

func subsets(nums []int) [][]int {
	var retVal [][]int
	var pri []int

	worker(pri, nums, &retVal)

	return retVal
}

func worker(prefix []int, next []int, res *[][]int) {
	dup := make([]int, len(prefix))
	copy(dup, prefix)
	*res = append(*res, dup)

	for i, val := range next {
		newPrefix := append(dup, val)
		worker(newPrefix, next[i+1:], res)
	}
}

func TestLC78(t *testing.T) {
	set := []int{1, 2, 3}
	set2 := []int{9, 0, 3, 1, 7}

	fmt.Println(subsets(set))
	fmt.Println(subsets(set2))
}
