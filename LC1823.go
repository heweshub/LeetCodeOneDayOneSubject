package main

func findTheWinner(n int, k int) int {
	m := make(map[int]bool, n+1)
	isLive := n
	kk := k
	i := 1
	for isLive > 1 {
		if i > n {
			i = 1
		}
		if !m[i] {
			kk--
			if kk == 0 {
				m[i] = true
				isLive--
				kk = k
			}
		}
		i++
	}
	for j := 1; j <= n; j++ {
		if _, ok := m[j]; !ok {
			return j
		}
	}
	return i
}
