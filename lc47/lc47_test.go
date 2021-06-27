package lc47

import (
	"fmt"
	"sort"
	"testing"
)

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var prefix []int
	sort.Ints(nums)
	worker(prefix, nums, &ans)
	return ans
}

func worker(prefix []int, nums []int, ans *[][]int) {
	if len(nums) == 1 {
		res := append(prefix, nums...)
		*ans = append(*ans, res)
		return
	}

	index := -1
	for i := range nums {
		if index != -1 && nums[i] == nums[index] {
			continue
		}

		index = i
		nextPrefix := append(prefix, nums[i])
		tmp := append([]int{}, nums[:i]...)
		nextNums := append(tmp, nums[i+1:]...)
		worker(nextPrefix, nextNums, ans)
	}
}

func TestLC47(t *testing.T) {
	nums := []int{1,2,1}

	fmt.Println(permuteUnique(nums))
}
