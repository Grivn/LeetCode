package main

import (
	"fmt"
)

func sortList(head *ListNode) *ListNode {
	return sort(head)
}

func sort(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	setLen := length(head)

	if setLen == 1 {
		return head
	}

	middle := setLen/2
	left, right := div(head, middle)

	return merge(sort(left), sort(right))
}

func length(head *ListNode) int {
	count := 0

	if head == nil {
		return count
	}

	ptr := head

	for {
		if ptr == nil {
			break
		}

		ptr = ptr.Next
		count++
	}

	return count
}

func div(head *ListNode, middle int) (*ListNode, *ListNode) {
	count := 0

	var pre, ptr *ListNode
	ptr = head

	for {
		if count == middle {
			pre.Next = nil
			break
		}

		count++
		pre = ptr
		ptr = ptr.Next
	}

	return head, ptr
}

func merge(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	var res, ptr *ListNode

	if left.Val < right.Val {
		res = left
		left = left.Next
	} else {
		res = right
		right = right.Next
	}

	ptr = res

	for {
		if left == nil || right == nil {
			break
		}

		if left.Val < right.Val {
			ptr.Next = left
			ptr = ptr.Next
			left = left.Next
		} else {
			ptr.Next = right
			ptr = ptr.Next
			right = right.Next
		}
	}

	if left == nil {
		ptr.Next = right
	}

	if right == nil {
		ptr.Next = left
	}

	return res
}

func main() {
	set := []int{4, 2, 1, 3}
	head := createList(set)
	res := sortList(head)
	fmt.Println("========================")
	println(length(res))
	printItem(res)

	//s1 := []int{1,3}
	//s2 := []int{2,4}
	//h1 := createList(s1)
	//h2 := createList(s2)
	//
	//res := merge(h1, h2)
	//printItem(res)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func createList(set []int) *ListNode {
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
