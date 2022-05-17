package main

//func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
//	// 中序遍历记录节点，然后再遍历一遍返回p的下一个节点
//	nodes := []*TreeNode{}
//	var inorder func(node *TreeNode)
//	inorder = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		inorder(node.Left)
//		nodes = append(nodes, node)
//		inorder(node.Right)
//	}
//	for i := 0; i < len(nodes); i++ {
//		if nodes[i] == p {
//			if i+1 < len(nodes) {
//				return nodes[i+1]
//			}
//		}
//	}
//	return nil
//}

// inorderSuccessor use two pointer to mark the p and the p's next TreeNode
//func inorderSuccessor(root, p *TreeNode) *TreeNode {
//	st := []*TreeNode{}
//	var pre, cur *TreeNode = nil, root
//	for len(st) > 0 || cur != nil {
//		// 将左子树都加入st中
//		for cur != nil {
//			st = append(st, cur)
//			cur = cur.Left
//		}
//		// 取出最后一个左子树，依次进行比较
//		cur = st[len(st)-1]
//		st = st[:len(st)-1]
//		if pre == p {
//			return cur
//		}
//		pre = cur
//		// 左中右，然后就指向右子树
//		cur = cur.Right
//	}
//	return nil
//}

// inorderSuccessor use the advantage of BST
func inorderSuccessor(root, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}
	node := root
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}
