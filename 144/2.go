package main

import "container/list"

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	cur := root
	l := list.New()
	l.Init()

	ret := []int{}

	for cur != nil || l.Len() != 0 {
		for cur != nil {
			ret = append(ret, cur.Val)
			l.PushFront(cur)
			cur = cur.Left
		}
		cur = l.Front().Value.(*TreeNode)
		l.Remove(l.Front())
		cur = cur.Right
	}

	return ret
}
