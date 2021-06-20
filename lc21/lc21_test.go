package lc21

import (
	"fmt"
	"testing"
)
/**
* Definition for singly-linked list.*/
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	//if l1 == nil && l2 == nil {
	//	return nil
	//}
	//
	//if l1 == nil && l2 != nil {
	//	return l2
	//}
	//
	//if l1 != nil && l2 == nil {
	//	return l1
	//}

	head := &ListNode{}
	pre := head
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil {
			pre.Next = l2
			break
		}

		if l2 == nil {
			pre.Next = l1
			break
		}

		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
			pre = pre.Next
			continue
		}

		if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
			pre = pre.Next
			continue
		}

		if l1.Val == l2.Val {
			pre.Next = l1
			l1 = l1.Next
			pre = pre.Next
			continue
		}
	}

	return head.Next
}

func TestLC21(t *testing.T) {
	l1 := generateList([]int{})
	l2 := generateList([]int{})

	printList(l1)
	printList(l2)


	printList(mergeTwoLists(l1, l2))
}

func generateList(set []int) *ListNode {
	if len(set) == 0 {
		return nil
	}

	head := &ListNode{Val: set[0], Next: nil}
	pre := head
	for i, num := range set {
		if i == 0 {
			continue
		}
		node := &ListNode{Val: num, Next: nil}
		pre.Next = node
		pre = node
	}
	return head
}

func printList(head *ListNode) {
	pre := head
	for {
		if pre == nil {
			break
		}
		fmt.Printf("%d ", pre.Val)
		if pre.Next == nil {
			break
		}
		pre = pre.Next
	}
	fmt.Printf("\n")
}
