package main

/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
1
\
2
/
3

Output: [1,3,2]
*/

func inorderTraversal(root *TreeNode) []int {

	return core(root)
}

func core(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := append(core(root.Left), root.Val)
	ret = append(ret, core(root.Right)...)
	return ret
}
