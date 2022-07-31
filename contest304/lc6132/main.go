package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumOperations([]int{0}))
}

func minimumOperations(nums []int) int {
	sort.Ints(nums)

	op := 0
	sub := 0
	for _, val := range nums {

		if val == 0 {
			continue
		}
		val -= sub

		if val > 0 {
			sub += val
			op++
		}
	}

	return op
}
