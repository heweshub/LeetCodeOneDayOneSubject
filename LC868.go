package main

import (
	"fmt"
	"strconv"
)

func binaryGap(n int) int {
	numStr := numTobin(n)
	if numStr == "0" {
		return 0
	}
	frontIdx := -1
	maxLen := 0
	for i, v := range numStr {
		if v == '1' {
			if frontIdx == -1 {
				frontIdx = i
			} else {
				maxLen = max11(maxLen, i-frontIdx)
			}
			frontIdx = i
		}
	}
	return maxLen
}
func numTobin(num int) string {
	s := ""
	for num != 0 {
		s = strconv.Itoa(num%2) + s
		num /= 2
	}
	return s
}

func max11(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(binaryGap(22))
}
