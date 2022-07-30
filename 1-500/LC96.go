package main

func numTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			g[i] += g[j-1] * g[i-j]
			// fmt.Printf("g[%d]=%d*g[%d]=%d=%d\n", j-1, g[j-1], i-j, g[i-j], g[i])
		}
	}
	return g[n]
}
