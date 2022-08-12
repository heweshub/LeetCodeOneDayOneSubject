package main

func preorderTraversal(root *TreeNode) (vals []int) {
	st := []*TreeNode{}
	node := root
	if node != nil || len(st) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			st = append(st, node)
			node = node.Left
		}
		node = st[len(st)-1].Right
		st = st[:len(st)-1]
	}
	return
}
