package main

func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return right
	}
	return left
}
