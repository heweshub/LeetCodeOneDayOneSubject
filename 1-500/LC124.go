package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)

		cur := node.Val + left + right

		maxSum = max(maxSum, cur)
		return node.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
