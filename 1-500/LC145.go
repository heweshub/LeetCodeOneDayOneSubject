package main

func postorderTraversal(root *TreeNode) (res []int) {
	st := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(st) > 0 {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}
		root = st[len(st)-1]
		st = st[:len(st)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			st = append(st, root)
			root = root.Right
		}
	}
	return
}
