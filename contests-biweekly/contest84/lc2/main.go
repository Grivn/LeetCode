package main

import "fmt"

func main() {
	nums := []int{4, 1, 3, 3}
	ret := countBadPairs(nums)
	fmt.Println(ret)
}

// timeout
func countBadPairs(nums []int) int64 {
	var ret int64
	worker(0, len(nums)-1, nums, &ret)
	return ret
}

func diff(index int, nums []int) int {
	return nums[index] - index
}

func worker(x, y int, nums []int, ret *int64) {
	// fmt.Println(x, ",", y)
	if x == y {
		return
	}

	if x+1 == y {
		if y-x != nums[y]-nums[x] {
			*ret++
		}
		return
	}
	if x > y {
		return
	}

	for i := x + 1; i <= y; i++ {
		if i-x != nums[i]-nums[x] {
			*ret++
		}
	}
	worker(x+1, y, nums, ret)
}
