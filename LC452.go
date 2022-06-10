package main

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if maxRight < p[1] {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}
