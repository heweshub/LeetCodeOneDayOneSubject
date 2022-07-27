package main

func maxArea(h []int) int {
	Max := 0
	l, r := 0, len(h)-1
	for l < r {
		cur := min(h[r], h[l]) * (r - l)
		if cur > Max {
			Max = cur
		}
		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}
	return Max
}
