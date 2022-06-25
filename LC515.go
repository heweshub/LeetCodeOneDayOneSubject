package main

func largestValues(root *TreeNode) (ans []int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := []*TreeNode{}
		values := []int{}
		for _, v := range q {
			values = append(values, v.Val)
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		q = tmp
		ans = append(ans, maxOfArray(values))
	}
	return
}
