package main

/*
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:

Input:
    2
   / \
  1   3
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6
Output: false
Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
             is 5 but its right child's value is 4.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//hint: The inorder traversal of the binary search tree is ordered
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var GlobalStack = list.New()
	var initedLastNum = false
	var lastNum int

	for root != nil || GlobalStack.Len() > 0 {
		for root != nil {
			GlobalStack.PushFront(root)
			root = root.Left
		}
		if GlobalStack.Len() > 0 {
			elem := GlobalStack.Front()
			root = elem.Value.(*TreeNode)
			GlobalStack.Remove(elem)
			if initedLastNum == false {
				initedLastNum = true
			} else {
				if lastNum >= root.Val {
					return false
				}
			}
			lastNum = root.Val
			root = root.Right
		}
	}
	return true
}
