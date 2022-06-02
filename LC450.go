package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	default:
		// 左右子树都存在
		successor := root.Right
		// 找右子树的最下面的左子树，successor放在原先key的节点位置，满足二叉搜索树的性质
		for successor.Left != nil {
			successor = successor.Left
		}
		// 然后在右子树中删除successor
		successor.Right = deleteNode(root.Right, successor.Val)
		// 左子树不变
		successor.Left = root.Left
		return successor
	}
	return root
}
