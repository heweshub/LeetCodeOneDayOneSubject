package main

import "math"

func isVaildBST(root *TreeNode) bool {
	return dfs1(root, math.MinInt64, math.MaxInt64)
}

func dfs1(root *TreeNode, lower, high int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= high {
		return false
	}
	return dfs1(root.Left, lower, root.Val) && dfs1(root.Right, root.Val, high)
}
