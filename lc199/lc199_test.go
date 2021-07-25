package lc199

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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	queue := list.New()
	queue.PushBack(root)

	for {
		length := queue.Len()

		if length == 0 {
			break
		}

		value := -1
		for {
			e := queue.Front()
			item := e.Value.(*TreeNode)
			queue.Remove(e)
			length--

			if item.Val != -1 {
				value = item.Val
			}

			if item.Left != nil {
				queue.PushBack(item.Left)
			}

			if item.Right != nil {
				queue.PushBack(item.Right)
			}

			if length == 0 {
				if value != -1 {
					res = append(res, value)
				}
				break
			}
		}
	}

	return res
}

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

func dfs(node *TreeNode) {
	if node == nil {
		fmt.Printf("-1 ")
		return
	}

	dfs(node.Left)
	fmt.Printf("%d ", node.Val)
	dfs(node.Right)
}

func bfs(node *TreeNode) {
	queue := list.New()
	queue.PushBack(node)

	for {
		length := queue.Len()

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
			}

			if item.Right != nil {
				queue.PushBack(item.Right)
			}

			if length == 0 {
				fmt.Printf(" | ")
				break
			}
		}
	}
}

func TestLC199(t *testing.T) {
	set := []int{1,2,3,-1,5,-1,4}

	root := createBTree(set)

	dfs(root)
	fmt.Printf("\n")
	bfs(root)
	fmt.Printf("\n")

	fmt.Println(rightSideView(root))
}
