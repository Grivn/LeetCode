package lc45

import "testing"

func jump(nums []int) int {
	jumps := 0
	curEnd := 0
	curFarthest := 0
	for i:=0; i<len(nums)-1; i++ {
		if curFarthest < i+nums[i] {
			curFarthest = i+nums[i]
		}
		if i == curEnd {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps
}

func TestLC43(t *testing.T) {

}
