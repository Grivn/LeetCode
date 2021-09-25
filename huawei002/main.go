package main

import (
	"fmt"
	"io"
	"sort"
)

func myWorker(remain []int, ave int) (bool, [][]int) {
	if len(remain) == 1 {
		return false, nil
	}

	if len(remain) == 2 {
		if checkMod(remain[0], remain[1], ave) {
			return true, [][]int{[]int{remain[1], remain[0]}}
		} else {
			return false, nil
		}
	}

	val1 := remain[0]

	remain = remain[1:]

	for i, val2 := range remain {
		if checkMod(val1, val2, ave) {
			next := append([]int{}, remain...)

			if i == len(remain)+1 {
				next = next[:i]
			} else if i == 0 {
				next = next[1:]
			} else {
				next = append(next[:i], next[i+1:]...)
			}

			success, partRes := myWorker(next, ave)

			if !success {
				return false, nil
			}

			part := []int{val2, val1}

			partRes = append(partRes, part)

			return true, partRes
		}
	}

	return false, nil
}

func checkMod(a, b, ave int) bool {
	if (a+b)%ave == 0 {
		return true
	}
	return false
}

type SortableOutput [][]int

func (s SortableOutput) Len() int {
	return len(s)
}

func (s SortableOutput) Less(i, j int) bool {
	sum1 := s[i][0] + s[i][1]
	sum2 := s[j][0] + s[j][1]

	if sum1 > sum2 {
		return true
	}

	if sum1 < sum2 {
		return false
	}

	return s[i][0]>s[j][0]
}

func (s SortableOutput) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var inputs []int

	var aveScore, n int

	_, err := fmt.Scan(&aveScore, &n)

	if err == io.EOF {
		return
	}

	for i:=0; i<n*2; i++ {
		var num int

		_, _ = fmt.Scan(&num)

		inputs = append(inputs, num)
	}

	sort.Ints(inputs)

	success, res := myWorker(inputs, aveScore)

	if !success {
		fmt.Println(0)
		return
	}

	var sortable SortableOutput

	sortable = append(sortable, res...)

	sort.Sort(sortable)

	var str string
	for _, part := range sortable {
		if str == "" {
			str = fmt.Sprintf("%d %d", part[0], part[1])
		} else {
			str = fmt.Sprintf("%s %d %d", str, part[0], part[1])
		}
	}
	str = fmt.Sprintf("%s\n", str)
	fmt.Printf(str)
}
