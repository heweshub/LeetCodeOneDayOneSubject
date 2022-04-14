package main

func maximumWealth(accounts [][]int) int {
	max := sum(accounts[0])
	for i := 1; i < len(accounts); i++ {
		item := sum(accounts[i])
		if item > max {
			max = item
		}
	}
	return max
}

func sum(a []int) (ans int) {
	for _, v := range a {
		ans += v
	}
	return
}
