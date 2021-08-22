package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		speed, err := input()
		if err == io.EOF {
			break
		}

		if len(speed) == 1 {
			fmt.Println(1)
			continue
		}

		sortedSpeeds := sort(speed)

		maxCount := 1

		for index := range sortedSpeeds {
			maxIndex := index + maxCount - 1

			if maxIndex+1 >= len(sortedSpeeds) {
				break
			}

			maxCount = maxInt(maxCount, legalCount(sortedSpeeds, index, maxCount))
		}

		fmt.Println(maxCount)
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func legalCount(sortedSpeeds []int, index int, maxCount int) int {

	if sortedSpeeds[index+maxCount-1] - sortedSpeeds[index] > 10 {
		return maxCount
	}

	count := maxCount

	for i:=index+maxCount; i<len(sortedSpeeds); i++ {
		if sortedSpeeds[i] <= sortedSpeeds[index]+10 {
			count++
		} else {
			break
		}
	}

	return count
}

func input() ([]int, error) {
	var n int
	_, err := fmt.Scan(&n)

	if err != nil {
		return nil, err
	}

	speeds := make([]int, n)

	for i:=0; i<n; i++ {
		_, err = fmt.Scan(&speeds[i])
		if err != nil {
			return nil, err
		}
	}

	return speeds, nil
}

func sort(nums []int) []int {
	length := len(nums)

	if length == 1 {
		return nums
	}

	middle := length/2
	left := nums[:middle]
	right := nums[middle:]

	return merge(sort(left), sort(right))
}

func merge(left, right []int) []int {

	var res []int

	for {
		if len(left) == 0 || len(right) == 0 {
			break
		}

		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	if len(left) == 0 {
		res = append(res, right...)
	}

	if len(right) == 0 {
		res = append(res, left...)
	}

	return res
}
