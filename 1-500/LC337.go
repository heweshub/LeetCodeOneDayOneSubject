package main

func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notselected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notselected}
}

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }
