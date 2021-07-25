package lc61

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func tail(head *ListNode) *ListNode {
	for {
		if head.Next == nil {
			return head
		}
		head = head.Next
	}
}

func length(head *ListNode) int {
	count := 0
	for {
		count++
		if head.Next == nil {
			break
		}

		head = head.Next
	}

	return count
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	k = k%length(head)
	count := length(head)-k

	if k == 0 {
		return head
	}

	tailItem := tail(head)

	oldFront := head

	pre := head
	item := pre
	for {
		if count == 0 {
			pre.Next = nil
			break
		}
		pre = item
		count--
		item = pre.Next
	}

	newFront := item

	tailItem.Next = oldFront

	head = newFront
	return head
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

func TestLC61(t *testing.T) {
	set := []int{1, 2, 3, 4, 5}
	h1 := generate(set)
	printItem(h1)
	println(length(h1))
	h2 := rotateRight(h1, 3)
	printItem(h2)
}
