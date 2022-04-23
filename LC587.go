package main

func outerTrees(trees [][]int) (ans [][]int) {
	n := len(trees)
	if n < 4 {
		return trees
	}
	leftMost := 0
	for i, tree := range trees {
		if tree[0] < trees[leftMost][0] {
			leftMost = i // 找最左边的点
		}
	}
	vis := make([]bool, n)
	p := leftMost
	for {
		q := (p + 1) % n
		for r, tree := range trees {
			// 大于0, r在q的右侧
			if cross(trees[p], trees[q], tree) < 0 {
				q = r // 找到最右边的点
			}
		}
		for i, b := range vis {
			if !b && i != p && i != q && cross(trees[p], trees[q], trees[i]) == 0 { // 在一条线上也需要加入栅栏中
				ans = append(ans, trees[i])
				vis[i] = true
			}
		}
		if !vis[q] {
			ans = append(ans, trees[q])
			vis[q] = true
		}
		p = q
		if p == leftMost { // 循环结束的条件，循环完相当于画了一圈
			break
		}
	}
	return
}

func cross(p, q, r []int) int {
	return (q[0]-p[0])*(r[1]-p[1]) - (q[1]-p[1])*(r[0]-q[0])
}
