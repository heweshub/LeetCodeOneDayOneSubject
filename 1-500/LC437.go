package main

func pathSum(root *TreeNode, target int) (ans int) {
	// 前缀和map
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, cur int64) {
		if node == nil {
			return
		}
		cur += int64(node.Val)
		ans += preSum[cur-int64(target)]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		preSum[cur]--
		return
	}
	dfs(root, 0)
	return
}

func rootSum(root *TreeNode, target int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == target {
		res++
	}
	res += rootSum(root.Left, target-val)
	res += rootSum(root.Right, target-val)
	return
}

func pathSum1(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, target)
	return res
}
