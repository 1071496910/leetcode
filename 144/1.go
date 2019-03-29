package main

/*Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
1
\
2
/
3

Output: [1,2,3]*/

func preorderTraversal(root *TreeNode) []int {
	return core(root)
}

func core(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := append([]int{root.Val}, core(root.Left)...)
	ret = append(ret, core(root.Right)...)
	return ret
}
