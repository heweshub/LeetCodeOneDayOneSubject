package main

import (
	"strconv"
)

func binaryGap(n int) (ans int) {
	numStr := numTobin(n)
	if numStr == "0" {
		return 0
	}
	frontIdx := -1
	for i, v := range numStr {
		if v == '1' {
			if frontIdx == -1 {
				frontIdx = i
			} else {
				ans = max11(ans, i-frontIdx)
			}
			frontIdx = i
		}
	}
	return
}
func numTobin(num int) (s string) {
	for num != 0 {
		s = strconv.Itoa(num%2) + s
		num /= 2
	}
	return
}

func max11(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// func main() {
// 	fmt.Println(binaryGap(22))
// }
