package lc206

import (
	"container/list"
	"fmt"
	"testing"
)
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	stack := list.New()

	temp := head
	for {
		if temp == nil {
			break
		}

		stack.PushBack(temp)
		temp = temp.Next
	}

	var newHead *ListNode
	ptr := newHead
	for {
		if stack.Len() == 0 {
			break
		}

		e := stack.Back()
		node := e.Value.(*ListNode)
		stack.Remove(e)

		if ptr == nil {
			newHead = node
			ptr = newHead
		} else {
			ptr.Next = node
			ptr = node
		}
	}

	if ptr != nil {
		ptr.Next = nil
	}

	return newHead
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

func TestLC206(t *testing.T) {
	set := []int{1,2,3,4,5}
	h1 := generate(set)
	h2 := reverseList(h1)
	printItem(h2)
}
