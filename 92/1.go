package main

/*
Reverse a linked list from position m to n. Do it in-place and in one-pass.

For example:
Given 1->2->3->4->5->NULL, m = 2 and n = 4,

return 1->4->3->2->5->NULL.

Note:
Given m, n satisfy the following condition:
1 ≤ m ≤ n ≤ length of list.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if m == 1 {
		return reverseN(head, (n - m))
	} else {
		tmphead := head
		for i := 0; i < m-2; i++ {
			tmphead = tmphead.Next
		}
		newList := reverseN(tmphead.Next, (n - m))
		tmphead.Next = newList
	}
	return head

}

func reverseN(head *ListNode, n int) *ListNode {
	newHead := head
	newNext := head
	for i := 0; i < n; i++ {
		tmpNext := newNext.Next
		newNext.Next = tmpNext.Next
		tmpNext.Next = newHead
		newHead = tmpNext
	}
	return newHead
}
