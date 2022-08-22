package main

func preorderTraversal(root *TreeNode) (vals []int) {
	st := []*TreeNode{}
	node := root
	// 先序遍历
	for node != nil || len(st) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			st = append(st, node)
			// 遍历完所有的左子树
			node = node.Left
		}
		// 切换到右子树进行遍历
		node = st[len(st)-1].Right
		st = st[:len(st)-1]
	}
	return
}
