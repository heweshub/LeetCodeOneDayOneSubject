package main

func verifyPostorder(p []int) bool {
	if len(p) < 1 {
		return true
	}
	root := p[len(p)-1]
	idx := -1
	for i, v := range p[:len(p)-1] {
		if v > root {
			idx = i
			break
		}
	}
	if idx == -1 {
		idx = len(p) - 1
	}
	left := p[:idx]
	right := p[idx : len(p)-1]
	// fmt.Println("left:", left)
	// fmt.Println("right:", right)
	for _, v := range left {
		if v > root {
			return false
		}
	}
	for _, v := range right {
		if v < root {
			return false
		}
	}
	return verifyPostorder(left) && verifyPostorder(right)
}
