package lc27

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		if nums[0] == val {
			nums = append(nums[:0])
			return 0
		} else {
			return 1
		}
	}

	if len(nums) == 0 {
		return 0
	}

	i := 0
	for {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}

		if len(nums) == i {
			break
		}
	}

	return len(nums)
}

func TestLC27(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2} // Input array
	fmt.Println(nums)
	k := removeElement(nums, 2) // Calls your implementation

	fmt.Println(k)
	fmt.Println(nums)
}
