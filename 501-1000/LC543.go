package main

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		ans = max(ans, l+r+1)
		return max(l, r) + 1
	}
	dfs(root)
	return ans - 1
}
