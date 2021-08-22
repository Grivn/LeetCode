package main

import (
	"flag"
	"fmt"
	"io"
)

func main() {
	for {
		inputs, err := input()
		if err == io.EOF {
			break
		}

		keyCount := len(inputs[0])

		for i:=0; i<keyCount; i++ {
			times := getTimes(inputs, i)
			calculate(times)
		}
	}
}

func calculate(times []int) float64 {
	total := float64(0)

	p := float64(1/len(times))

	for _, t := range times {
		partial := p*float64(t)


	}

	return 0
}

func getTimes(inputs [][]int, keyIndex int) []int {
	var res []int

	for _, keys := range inputs {
		res = append(res, keys[keyIndex])
	}
	return sort(res)
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

func input() ([][]int, error) {
	var n, m int

	_, err := fmt.Scan(&n, &m)

	if err != nil {
		return nil, err
	}

	res := make([][]int, n)

	for i:=0; i<n; i++ {
		keys := make([]int, m)

		for j:=0; j<m; j++ {
			_, _ = fmt.Scan(&keys[j])
		}
		res[i] = keys
	}

	return res, nil
}
