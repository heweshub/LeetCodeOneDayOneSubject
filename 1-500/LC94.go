package main

func inorderTraversal(root *TreeNode) (res []int) {
	st := []*TreeNode{}
	for root != nil || len(st) > 0 {
		if root != nil {
			st = append(st, root)
			root = root.Left
		}
		root = st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
