package lc55

import "testing"

func canJump(nums []int) bool {
	curFarthest := 0
	curEnd := 0

	for i:=0; i<len(nums)-1; i++ {
		if curFarthest < i+nums[i] {
			curFarthest = i+nums[i]
		}
		if i == curEnd {
			curEnd = curFarthest
		}
	}

	return curEnd>=len(nums)-1
}

func TestLC55(t *testing.T) {
	nums := []int{3,2,1,0,4}
	println(canJump(nums))
}
