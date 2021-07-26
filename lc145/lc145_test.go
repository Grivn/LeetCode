package lc145

import (
	"container/list"
	"fmt"
	"testing"
)
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
			filledLeft = true
			if node.Val == -1000 {
				continue
			}
			e := queue.Front()
			item := e.Value.(*TreeNode)
			item.Left = node
			if val != -1000 {
				queue.PushBack(node)
			}
		} else {
			filledLeft = false
			if node.Val == -1000 {
				continue
			}
			e := queue.Front()
			item := e.Value.(*TreeNode)
			item.Right = node
			queue.Remove(e)
			if val != -1000 {
				queue.PushBack(node)
			}
		}
	}

	return root
}

func bfs(node *TreeNode) {
	queue := list.New()
	queue.PushBack(node)

	hasNext := true
	for {
		length := queue.Len()
		if !hasNext {
			break
		}
		hasNext = false
		if length == 0 {
			break
		}

		for {
			e := queue.Front()
			item := e.Value.(*TreeNode)
			queue.Remove(e)
			length--
			fmt.Printf("%d ", item.Val)

			if item.Left != nil {
				queue.PushBack(item.Left)
				hasNext = true
			} else {
				NilNode := &TreeNode{Val: -1000}
				queue.PushBack(NilNode)
			}

			if item.Right != nil {
				queue.PushBack(item.Right)
				hasNext = true
			} else {
				NilNode := &TreeNode{Val: -1000}
				queue.PushBack(NilNode)
			}

			if length == 0 {
				fmt.Printf(" | ")
				break
			}
		}
	}
}
func TestLC145(t *testing.T) {

}
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	if root.Left != nil {
		lRes := postorderTraversal(root.Left)
		res = append(res, lRes...)
	}

	if root.Right != nil {
		rRes := postorderTraversal(root.Right)
		res = append(res, rRes...)
	}

	res = append(res, root.Val)

	return res
}
