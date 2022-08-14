package main

import "sort"

func main() {

}
func edgeScore(edges []int) int {
	record := make(map[int]int)
	largestId := make([]int, 0)
	max := 0
	for i, val := range edges {
		record[val] += i
		if record[val] > max {
			max = record[val]
			largestId = make([]int, 0)
			largestId = append(largestId, val)
			continue
		}
		if record[val] == max {
			largestId = append(largestId, val)
			continue
		}
	}
	sort.Ints(largestId)
	if len(largestId) == 0 {
		return 0
	}
	return largestId[0]
}
