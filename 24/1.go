package main

/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead = new(ListNode)
	var tmpHead = new(ListNode)
	var tmp = new(ListNode)

	newHead.Next = head
	tmpHead = newHead

	for head != nil && head.Next != nil {
		tmp = head.Next
		head.Next = tmp.Next
		tmp.Next = head
		tmpHead.Next = tmp

		head = head.Next
		tmpHead = tmp.Next
	}

	return newHead.Next
}
