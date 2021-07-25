package lc116and117

import (
	"container/list"
	"fmt"
	"testing"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func createBTree(set []int) *Node {
	root := new(Node)
	queue := list.New()
	leftFilled := false

	for _, val := range set {
		node := &Node{Val: val}

		if queue.Len() == 0 {
			root = node
			queue.PushBack(node)
		} else if !leftFilled {
			leftFilled = true
			e := queue.Front()
			item := e.Value.(*Node)
			item.Left = node

			if val != -1 {
				queue.PushBack(node)
			}
		} else {
			leftFilled = false
			e := queue.Front()
			item := e.Value.(*Node)
			item.Right = node
			queue.Remove(e)

			if val != -1 {
				queue.PushBack(node)
			}
		}
	}

	return root
}

func bfs(root *Node) {
	queue := list.New()
	queue.PushBack(root)

	for {
		length := queue.Len()

		if length == 0 {
			break
		}

		for {
			e := queue.Front()
			item := e.Value.(*Node)
			queue.Remove(e)
			length--

			fmt.Printf("%d ", item.Val)
			if item.Next == nil {
				fmt.Printf(" stop ")
			}

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

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := list.New()
	queue.PushBack(root)

	for {
		length := queue.Len()

		if length == 0 {
			break
		}

		var prefix *Node
		for {
			e := queue.Front()
			item := e.Value.(*Node)
			queue.Remove(e)
			length--

			if prefix == nil {
				prefix = item
			} else {
				prefix.Next = item
				prefix = item
			}

			if item.Left != nil {
				queue.PushBack(item.Left)
			}

			if item.Right != nil {
				queue.PushBack(item.Right)
			}

			if length == 0 {
				break
			}
		}
	}
	return root
}

func TestLC16(t *testing.T) {
	set := []int{1,2,3,4,5,6,7}
	root := createBTree(set)
	//bfs(root)
	connect(root)
	bfs(root)
}
