package main

/*
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	l := head
	r := head
	len := 0
	//k great than len
	kgtl := false
	var newHead *ListNode

	for i := 0; i < k; i++ {
		r = r.Next
		len++
		if r == nil {
			kgtl = true
			break
		}
	}

	if kgtl {
		k = k % len
		if k == 0 {
			return head
		}
		r = head
		for i := 0; i < k; i++ {
			r = r.Next
		}
	}

	for r.Next != nil {
		l = l.Next
		r = r.Next
	}
	newHead = l.Next
	r.Next = head
	l.Next = nil
	return newHead
}
