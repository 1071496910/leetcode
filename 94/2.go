package main

import "container/list"

func inorderTraversal2(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return []int{}
	}

	l := list.New()
	l.Init()
	l.PushFront(root)

	cur := root
	for l.Len() != 0 || cur != nil {
		for cur != nil {
			l.PushFront(cur)
			cur = cur.Left
		}
		cur = l.Front().Value.(*TreeNode)
		l.Remove(l.Front())
		ret = append(ret, cur.Val)
		cur = cur.Right
	}

	return ret
}
