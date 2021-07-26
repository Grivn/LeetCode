package lc95

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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return worker(1, n)
}

func worker(start int, stop int) []*TreeNode {
	if start == 0 {
		return nil
	}

	if start == stop {
		node := &TreeNode{Val: start}
		return []*TreeNode{node}
	}

	var res []*TreeNode
	for index:=start; index<=stop; index++ {
		if index == start {
			rNodes := worker(start+1, stop)
			for _, node := range rNodes {
				cur := &TreeNode{Val: index, Right: node}
				fmt.Printf("[%d -> %d] cur %v rNode %v\n", start, stop, cur, node)
				res = append(res, cur)
			}
			continue
		}

		if index == stop {
			lNodes := worker(start, index-1)
			for _, node := range lNodes {
				cur := &TreeNode{Val: index, Left: node}
				fmt.Printf("[%d -> %d] cur %v lNode %v\n", start, stop, cur, node)
				res = append(res, cur)
			}
			continue
		}

		lNodes := worker(start, index-1)
		rNodes := worker(index+1, stop)
		for _, lNode := range lNodes {
			for _, rNode := range rNodes {
				cur := &TreeNode{Val: index, Left: lNode, Right: rNode}
				fmt.Printf("[%d -> %d] cur %v lNode %v rNode %v\n", start, stop, cur, lNode, rNode)
				res = append(res, cur)
			}
		}
	}

	return res
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

			if item.Right == nil && item.Left == nil {
				if length == 0 {
					//fmt.Printf(" | ")
					break
				}
				continue
			}

			if item.Left != nil {
				queue.PushBack(item.Left)
			} else {
				fmt.Printf("-1 ")
			}

			if item.Right != nil {
				queue.PushBack(item.Right)
			} else {
				fmt.Printf("-1 ")
			}

			if length == 0 {
				//fmt.Printf(" | ")
				break
			}
		}
	}
}

func TestLC95(t *testing.T) {
	roots := generateTrees(3)

	for _, root := range roots {
		bfs(root)
		fmt.Printf("\n")
	}
}
