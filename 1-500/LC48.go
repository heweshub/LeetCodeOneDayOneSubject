package main

func rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := 0; i < (n+1)/2; j++ {
			m[i][j], m[n-1-j][i], m[n-1-j][n-1-i], m[j][n-1-i] = m[n-1-j][i], m[n-1-j][n-1-i], m[j][n-1-i], m[i][j]
		}
	}
}
