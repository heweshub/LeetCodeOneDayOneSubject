package main

func isUnivalTree(root *TreeNode) bool {
	value := root.Val
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != value {
			return false
		}
		return dfs(node.Left) && dfs(node.Right)
	}
	return dfs(root.Left) && dfs(root.Right)
}
