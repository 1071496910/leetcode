package main

import (
	"fmt"
)

/*
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func genTrees(left int, right int) []*TreeNode {
	if left > right {
		return nil
	}
	if right == left {
		return []*TreeNode{
			&TreeNode{
				Val: left,
			},
		}
	}
	res := []*TreeNode{}
	for i := left; i <= right; i++ {
		leftSubTrees := genTrees(left, i-1)
		rightSubTrees := genTrees(i+1, right)
		for _, lt := range leftSubTrees {
			for _, rt := range rightSubTrees {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  lt,
					Right: rt,
				})
			}
		}
		if leftSubTrees == nil {
			for _, rt := range rightSubTrees {
				res = append(res, &TreeNode{
					Val:   i,
					Right: rt,
				})
			}
		}
		if rightSubTrees == nil {
			for _, lt := range leftSubTrees {
				res = append(res, &TreeNode{
					Val:  i,
					Left: lt,
				})
			}
		}
	}
	return res
}

func generateTrees(n int) []*TreeNode {
	return genTrees(1, n)
}
