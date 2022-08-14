package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2, 2, 3}
	set := frequencySort(nums)
	fmt.Println(set)
}

func frequencySort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return nums
	}
	m := make(map[int]*tuple)

	for _, num := range nums {
		t := m[num]
		if t == nil {
			m[num] = &tuple{val: num, count: 0}
			t = m[num]
		}
		t.count++
	}

	set := make(tupleSet, 0)
	for _, t := range m {
		set = append(set, *t)
	}
	//fmt.Println(set)
	//fmt.Println(m)
	sort.Sort(set)

	var res []int

	for _, t := range set {
		for {
			res = append(res, t.val)
			t.count--
			if t.count == 0 {
				break
			}
		}
	}
	return res
}

type tuple struct {
	val   int
	count int
}

type tupleSet []tuple

func (set tupleSet) Len() int      { return len(set) }
func (set tupleSet) Swap(i, j int) { set[i], set[j] = set[j], set[i] }
func (set tupleSet) Less(i, j int) bool {
	if set[i].count == set[j].count {
		return set[i].val > set[j].val
	}
	return set[i].count < set[j].count
}
