package lc101

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

func isSymmetric(root *TreeNode) bool {
	// bfs

	queue := list.New()
	queue.PushBack(root)
	layer := 0

	hasNext := true
	for {
		layer++
		length := queue.Len()

		if !hasNext {
			break
		}

		if length == 0 {
			break
		}

		stack := list.New()
		count := 0
		div := length/2
		hasNext = false
		for {
			e := queue.Front()
			item := e.Value.(*TreeNode)
			queue.Remove(e)
			count++
			length--

			// push stack
			if layer == 1 {
				// do nothing
			} else {
				if count <= div {
					stack.PushBack(item)
					fmt.Printf("push %v\n", item)
				} else {
					back := stack.Back()
					backItem := back.Value.(*TreeNode)
					if backItem.Val == item.Val {
						stack.Remove(back)
						fmt.Printf("pop %v\n", back.Value.(*TreeNode))
					} else {
						return false
					}
				}
			}

			if item.Val == -1000 {
				// do nothing
			} else {
				fmt.Printf("item %v l %v r %v\n", item, item.Left, item.Right)
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
			}

			if length == 0 {
				fmt.Printf("has next %v\n", hasNext)
				break
			}
		}
	}
	return true
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
			if val != -1000 {
				queue.PushBack(node)
			}
		} else {
			e := queue.Front()
			item := e.Value.(*TreeNode)
			item.Right = node
			filledLeft = false
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
func TestLC101(t *testing.T) {
	set := []int{9,-42,-42,-1000,76,76,-1000,-1000,13,-1000,13}
	root := createBTree(set)
	bfs(root)
	fmt.Println()
	println(isSymmetric(root))
}
