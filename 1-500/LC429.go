package main

// Definition for a Node.
// type Node struct {
// 	Val      int
// 	Children []*Node
// }

// 类似于二叉树的层次遍历
func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*Node{root}
	for len(q) > 0 {
		level := []int{}
		// tmp := q
		q = nil
		// for i := range tmp {
		// 	level = append(level, tmp[i].Val)
		// 	q = append(q, tmp[i].Children...)
		// }
		//for _, node := range tmp {
		//	level = append(level, node.Val)
		//	q = append(q, node.Children...)
		//}
		ans = append(ans, level)
	}
	return
}
