package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func countPrimeSetBits(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		ss := ten2er(i)
		cnt1 := strings.Count(ss, "1")
		if isZhishu(cnt1) {
			ans++
		}
	}
	return ans
}
func ten2er(x int) string {
	s := ""
	for x > 0 {
		s = strconv.Itoa(x%2) + s
		x /= 2
	}
	return s
}

func isZhishu(x int) bool {
	if x == 1 {
		return false
	}
	if x <= 3 {
		return true
	}
	n := int(math.Sqrt(float64(x)))
	for i := 2; i < n; i++ {
		if x%i == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}
