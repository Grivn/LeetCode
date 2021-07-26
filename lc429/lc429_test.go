package lc429

import (
	"container/list"
)

type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int

	queue := list.New()
	queue.PushBack(root)

	for {
		length := queue.Len()

		if length == 0 {
			break
		}

		var pRes []int
		for {
			e := queue.Front()
			item := e.Value.(*Node)
			queue.Remove(e)
			length--

			pRes = append(pRes, item.Val)

			for _, child := range item.Children {
				if child != nil {
					queue.PushBack(child)
				}
			}

			if length == 0 {
				res = append(res, pRes)
				break
			}
		}
	}

	return res
}
