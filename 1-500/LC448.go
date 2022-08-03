package main

func findDisappearedNumbers(a []int) (ans []int) {
	n := len(a)
	for _, v := range a {
		v = (v - 1) % n
		a[v] += n
	}
	for i, v := range a {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}
