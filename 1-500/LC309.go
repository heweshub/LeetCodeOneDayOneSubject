package main

func maxProfit(p []int) int {
	if len(p) == 0 {
		return 0
	}
	n := len(p)
	f := make([][3]int, n)
	f[0][0] = -p[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-p[i])
		f[i][1] = f[i-1][0] + p[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}
