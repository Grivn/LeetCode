package lc92

import (
	"container/list"
	"fmt"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		//lNode *ListNode
		rNode *ListNode

		prefix *ListNode
		behind *ListNode
	)

	count := 0

	temp := head
	stack := list.New()
	for {
		count++

		if count < left {
			prefix = temp
		}

		if count == left {
			//lNode = temp
		}

		if count == right {
			rNode = temp
			behind = temp.Next
			break
		}

		if count >= left && count < right {
			stack.PushBack(temp)
		}
		temp = temp.Next
	}

	newPrefix := rNode
	for {
		if stack.Len() == 0 {
			break
		}

		e := stack.Back()
		node := e.Value.(*ListNode)
		stack.Remove(e)

		newPrefix.Next = node
		newPrefix = node
	}

	newPrefix.Next = behind

	if prefix != nil {
		prefix.Next = rNode
		return head
	} else {
		return rNode
	}
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

func TestLC92(t *testing.T) {
	set := []int{3,5}
	h1 := generate(set)

	printItem(h1)

	h2 := reverseBetween(h1, 1, 2)

	printItem(h2)

	set2 := []int{1,2,3,4,5}
	h3 := generate(set2)

	h4 := reverseBetween(h3, 2, 4)
	printItem(h4)
}
