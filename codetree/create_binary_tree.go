package codetree

import "container/list"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// createBTree creates binary tree with layer order
func createBTree(set []int) *TreeNode {
	root := new(TreeNode)

	queue := list.New()
	filledLeft := false
	for _, val := range set {
		node := &TreeNode{Val: val}

		if queue.Len() == 0 {
			queue.PushBack(node)
			root = node
		} else if !filledLeft {
			e := queue.Front()
			item := e.Value.(*TreeNode)
			item.Left = node
			filledLeft = true
			if val != -1 {
				queue.PushBack(node)
			}
		} else {
			e := queue.Front()
			item := e.Value.(*TreeNode)
			item.Right = node
			filledLeft = false
			queue.Remove(e)
			if val != -1 {
				queue.PushBack(node)
			}
		}
	}

	return root
}
