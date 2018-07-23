package main

import (
	"container/list"
)

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

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue := list.New()
	var levelEnd *TreeNode
	var nextLevelEnd *TreeNode

	if root == nil {
		return res
	}

	tmpArr := []int{}
	levelEnd = root

	queue.PushBack(root)

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		tmpNode := elem.Value.(*TreeNode)
		tmpArr = append(tmpArr, tmpNode.Val)
		if tmpNode.Left != nil {
			queue.PushBack(tmpNode.Left)
			nextLevelEnd = tmpNode.Left
		}
		if tmpNode.Right != nil {
			queue.PushBack(tmpNode.Right)
			nextLevelEnd = tmpNode.Right
		}
		if tmpNode == levelEnd {
			res = append(res, tmpArr)
			tmpArr = []int{}
			levelEnd = nextLevelEnd
		}

	}
	return res

}

//[3,9,20,null,null,15,7]
func main() {
	tmpTree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 17,
			},
		},
	}

	levelOrder(tmpTree)

}
