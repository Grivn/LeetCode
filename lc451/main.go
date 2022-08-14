package main

import "sort"

func main() {

}
func frequencySort(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	m := make(map[rune]*tuple)

	for _, c := range s {
		t := m[c]
		if t == nil {
			m[c] = &tuple{val: c, count: 0}
			t = m[c]
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

	var res string

	for _, t := range set {
		for {
			res += string(t.val)
			t.count--
			if t.count == 0 {
				break
			}
		}
	}
	return res
}

type tuple struct {
	val   rune
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
