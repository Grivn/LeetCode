package lc234

import (
	"container/list"
	"fmt"
	"testing"
)

func isPalindrome(head *ListNode) bool {
	stack := list.New()

	temp := head
	for {
		if temp == nil {
			break
		}
		stack.PushBack(temp)
		temp = temp.Next
	}

	for {
		if stack.Len() == 0 {
			break
		}

		e := stack.Back()
		node := e.Value.(*ListNode)
		stack.Remove(e)

		if node.Val == head.Val {
			head = head.Next
			continue
		} else {
			return false
		}
	}

	return true
}

type ListNode struct {
	Val int
	Next *ListNode
}

func generate(set []int) *ListNode {
	if len(set) == 0 {
		return nil
	}

	head := &ListNode{Val: set[0], Next: nil}

	pre := head

	for index, i := range set {
		if index == 0 {
			continue
		}
		item := &ListNode{
			Val: i,
			Next: nil,
		}

		pre.Next = item
		pre = item
	}

	return head
}

func printItem(head *ListNode) {
	for {
		fmt.Printf("%d ", head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	fmt.Printf("\n")
}
func TestLC234(t *testing.T) {

}
