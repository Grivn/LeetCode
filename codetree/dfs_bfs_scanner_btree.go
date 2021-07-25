package codetree

import (
	"container/list"
	"fmt"
)

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
