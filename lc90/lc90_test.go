package lc90

import (
	"fmt"
	"sort"
	"testing"
)

func subsetsWithDup(nums []int) [][]int {
	var prefix []int
	var res [][]int
	sort.Ints(nums)
	worker(prefix, nums, &res)
	return res
}

func worker(prefix []int, next []int, res *[][]int) {
	fmt.Printf("prefix %v, next %v\n", prefix, next)
	dup := make([]int, len(prefix))
	copy(dup, prefix)
	*res = append(*res, dup)

	filter := make(map[int]bool)
	for i, val := range next {
		if filter[val] {
			continue
		}

		filter[val] = true
		newPrefix := append(dup, val)
		worker(newPrefix, next[i+1:], res)
	}
}

func TestLC90(t *testing.T) {
	set := []int{1, 4, 4, 4}
	fmt.Println(subsetsWithDup(set))
}
