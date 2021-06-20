package lc26

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	count := 0
	i := 0
	for {
//		fmt.Printf("count: %d, i: %d, num: %d\n", count, i, nums[i])
		if i == 0 {
			count++
			i++
			continue
		}
		if nums[i] == nums[i-1] {
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			count++
			i++
		}
		if len(nums) == i {
			break
		}
	}
	return count
}

func TestLC26(t *testing.T) {
	nums := []int{1} // Input array
	fmt.Println(nums)
	k := removeDuplicates(nums) // Calls your implementation

	fmt.Println(k)
	fmt.Println(nums)
}
