package main

import "sort"

func main() {

}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	var set tupleSet
	for key, val := range m {
		t := tuple{val: key, count: val}
		set = append(set, t)
	}
	sort.Sort(set)

	var res []int
	for i, t := range set {
		res = append(res, t.val)
		if i == k-1 {
			break
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
	return set[i].count > set[j].count
}
