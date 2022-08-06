package main

import "sort"

func main() {

}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var res slice
	visited := make(map[int]bool)
	for _, sub1 := range items1 {
		key := sub1[0]
		visited[key] = true
		found := false
		for _, sub2 := range items2 {
			if sub2[0] == key {
				sub := []int{key, sub1[1] + sub2[1]}
				res = append(res, sub)
				found = true
				break
			}
		}
		if !found {
			sub := sub1
			res = append(res, sub)
		}
	}

	for _, sub2 := range items2 {
		if !visited[sub2[0]] {
			res = append(res, sub2)
		}
	}
	sort.Sort(res)
	return res
}

type slice [][]int

func (s slice) Len() int { return len(s) }

func (s slice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s slice) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
