package lc46

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {
	var prefix []int
	var ans [][]int
	subSet(prefix, nums, &ans)
	return ans
}

func subSet(prefix []int, nums []int, ans *[][]int) {
	if len(nums) == 1 {
		res := append(prefix, nums...)
		*ans = append(*ans, res)
		return
	}

	for i:=0; i<len(nums); i++ {
		nextPrefix := append(prefix, nums[i])
		tmp := append([]int{}, nums[:i]...)
		set := append(tmp, nums[i+1:]...)
		subSet(nextPrefix, set, ans)
	}
}

func TestLC46(t *testing.T) {

	nums := []int{1,2,3}
	fmt.Println(permute(nums))
}
