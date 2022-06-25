package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ls1, ls2 := addTree(root1), addTree(root2)
	// fmt.Println(ls1)
	// fmt.Println(ls2)
	return merge(ls1, ls2)
}

func merge(ls1, ls2 []int) (res []int) {
	n, m := len(ls1), len(ls2)
	i, j := 0, 0
	for i < n && j < m {
		if ls1[i] <= ls2[j] {
			res = append(res, ls1[i])
			i++
		} else {
			res = append(res, ls2[j])
			j++
		}
	}
	if i < n {
		res = append(res, ls1[i:]...)
	}
	if j < m {
		res = append(res, ls2[j:]...)
	}
	return
}

func addTree(node *TreeNode) (ans []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(node)
	return
}
