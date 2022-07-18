package main

func kthLargest(root *TreeNode, k int) (ans int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		k--
		if k == 0 {
			ans = node.Val
		}
		dfs(node.Left)
	}
	dfs(root)
	return
}
