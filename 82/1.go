package main

/*
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Example 1:

Input: 1->2->3->3->4->4->5
Output: 1->2->5
Example 2:

Input: 1->1->1->2->3
Output: 2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead = new(ListNode)
	var newtail = newHead

	l := head
	r := head.Next
	duplicated := false

	for r != nil {
		duplicated = false
		for r != nil && l.Val == r.Val {
			duplicated = true
			r = r.Next
		}
		if r == nil {
			break
		}
		if !duplicated {
			newtail.Next = l
			newtail = l
		}
		l = r
		r = r.Next
	}
	if newtail != nil {
		if l.Next == r {
			newtail.Next = l
			newtail = newtail.Next
		}
		newtail.Next = nil
	}

	return newHead.Next

}
