package main

func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return false
	}
	return recur(a, b) || isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}

func recur(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}
