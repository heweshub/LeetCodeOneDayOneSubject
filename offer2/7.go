package main

func buildTree(pre []int, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{pre[0], nil, nil}
	i := 0
	for ; i < len(in); i++ {
		if in[i] == pre[0] {
			break
		}
	}
	root.Left = buildTree(pre[1:len(in[:i])+1], in[:i])
	root.Right = buildTree(pre[len(in[:i])+1:], in[i+1:])
	return root
}
