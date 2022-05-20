package main

func largestCombination(candidates []int) (ans int) {
	for i := 0; i < 24; i++ {
		tmp := 0
		for _, v := range candidates {
			tmp += v >> i & 1
		}
		ans = Max(ans, tmp)
	}
	return
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	largestCombination([]int{16, 17, 71, 62, 12, 24, 14})
}
