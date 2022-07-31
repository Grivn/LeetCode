package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	/*
		[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
		30
	*/
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 30))
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sub := []int{}
	sort.Ints(candidates)
	helper(candidates, target, &result, sub)
	return result
}

func helper(candidates []int, target int, result *[][]int, sub []int) {

	if len(candidates) == 0 {
		return
	}
	if candidates[0] == target {
		sub = append(sub, candidates[0])
		*result = append(*result, sub)
		return
	} else if candidates[0] < target {
		for i := 0; i < len(candidates)-1; i++ {
			if candidates[i+1] != candidates[i] {
				helper(candidates[i+1:], target, result, sub)
				sub2 := make([]int, len(sub))
				copy(sub2, sub)
				sub2 = append(sub2, candidates[0])
				helper(candidates[1:], target-candidates[0], result, sub2)
				return
			}
		}
		sub3 := make([]int, len(sub)) //the case when all the rest numbers are the same
		copy(sub3, sub)
		sub3 = append(sub3, candidates[0])
		helper(candidates[1:], target-candidates[0], result, sub3)
	} else {
		return
	}
}

func worker(prefix []int, candidates []int, res *[][]int, target int) {
	pSet := make([]int, len(prefix))
	copy(pSet, prefix)

	sum := calSum(pSet)
	//fmt.Println("prefix ", pSet, " candidate ", candidates, " sum ", sum)
	if sum == target {
		sort.Ints(pSet)
		if !isExist(pSet, *res) {
			*res = append(*res, pSet)
			fmt.Println(res)
		}

		return
	}

	if sum > target {
		return
	}

	for i, val := range candidates {
		newPrefix := append(pSet, val)
		//fmt.Println("new prefix ", newPrefix, " candidate ", candidates)
		if i == len(candidates)-1 {
			worker(newPrefix, []int{}, res, target)
			break
		}
		newCandidate := candidates[i+1:]
		worker(newPrefix, newCandidate, res, target)
	}
}

func calSum(set []int) int {
	sum := 0

	for _, val := range set {
		sum += val
	}

	return sum
}

func isExist(set []int, res [][]int) bool {
	for _, subSet := range res {
		if reflect.DeepEqual(set, subSet) {
			return true
		}
	}
	return false
}
