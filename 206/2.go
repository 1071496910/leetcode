package main

/*
Reverse a singly linked list.

click to show more hints.

Hint:
A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// iterative version
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := head
	newTail := head

	for newTail != nil && newTail.Next != nil {
		tmpNext := newTail.Next
		newTail.Next = tmpNext.Next
		tmpNext.Next = newHead
		newHead = tmpNext
	}

	return newHead

}
