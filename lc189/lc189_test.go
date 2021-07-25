package lc189

import (
	"fmt"
	"testing"
)

func rotate(nums []int, k int) {
	head := generate(nums)
	newHead := rotateRight(head, k)

	newNums := rollBack(newHead)
	copy(nums, newNums)
}

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

func rollBack(head *ListNode) []int {
	var ret []int
	for {
		if head == nil {
			break
		}

		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}

func TestLC189(t *testing.T) {
	set := []int{1, 2, 3, 4, 5}
	rotate(set, 2)
	fmt.Println(set)
}
