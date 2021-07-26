package lc102

import "container/list"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
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
			item := e.Value.(*TreeNode)
			queue.Remove(e)
			length--

			pRes = append(pRes, item.Val)

			if item.Left != nil {
				queue.PushBack(item.Left)
			}

			if item.Right != nil {
				queue.PushBack(item.Right)
			}

			if length == 0 {
				res = append(res, pRes)
				break
			}
		}
	}
	return res
}