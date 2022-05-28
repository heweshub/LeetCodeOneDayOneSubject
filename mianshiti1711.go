package main

import "math"

func findClosest(words []string, word1 string, word2 string) (ans int) {
	ans = math.MaxInt32
	id1, id2 := -1, -1
	for i := range words {
		if words[i] == word1 {
			id1 = i
		} else if words[i] == word2 {
			id2 = i
		}
		if id1 != -1 && id2 != -1 {
			ans = min(ans, abs(id1, id2))
		}
	}
	return
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
