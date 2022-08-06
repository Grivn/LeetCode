package main

import "fmt"

func main() {
	/*
		   	[9,8,7,0,5,6,1,3,2,2]
		      1
		      6

		[4,4,8,-1,9,8,4,4,1,1]
		5
		6
	*/

	//list := []int{9, 8, 7, 0, 5, 6, 1, 3, 2, 2}
	list2 := []int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}
	fmt.Println(closestMeetingNode(list2, 5, 6))
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	ptr1 := node1
	ptr2 := node2

	var jumpList1 []int
	var jumpList2 []int

	jumpList1 = append(jumpList1, ptr1)
	scanned1 := make(map[int]bool)
	for {
		scanned1[ptr1] = true
		jump1 := edges[ptr1]
		ptr1 = jump1
		if jump1 == -1 || scanned1[jump1] {
			break
		}
		jumpList1 = append(jumpList1, jump1)
	}

	jumpList2 = append(jumpList2, ptr2)
	scanned2 := make(map[int]bool)
	for {
		scanned2[ptr2] = true
		jump2 := edges[ptr2]
		ptr2 = jump2
		if jump2 == -1 || scanned2[jump2] {
			break
		}
		jumpList2 = append(jumpList2, jump2)
	}

	ret := -1
	minDst := -1
	//fmt.Println(jumpList1)
	//fmt.Println(jumpList2)
	for dst1, val1 := range jumpList1 {
		for dst2, val2 := range jumpList2 {
			maxOne := dst1
			if dst1 < dst2 {
				maxOne = dst2
			}
			if val1 == val2 {
				sumDist := dst1 + dst2
				if minDst == -1 {
					minDst = maxOne
					ret = val1
				} else if minDst > sumDist {
					minDst = maxOne
					ret = val1
				} else if minDst == maxOne {
					if val1 < ret {
						ret = val1
					}
				}
			}
		}
	}

	return ret
}
