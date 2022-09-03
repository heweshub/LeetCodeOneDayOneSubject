package main

import (
	"fmt"
	"math"
)

var m = map[int]int{1: 0, 2: 1, 3: 2}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	sum := 0
	for _, v := range nums {
		sum += cal(v)
	}
	fmt.Println(sum)
}

func cal(v int) int {
	if _, ok := m[v]; ok {
		return m[v]
	}
	if isPrime(v) {
		return 1 + cal(v-1)
	} else {
		mid := int(math.Sqrt(float64(v)))
		var n1, n2 int
		tmp := math.MaxInt64
		for i := mid; i > 1; i-- {
			if v%i == 0 {
				n1, n2 = i, v/i
				tmp = min(tmp, cal(n1)+cal(n2)+1)
			}
		}
		return tmp
	}
}

func isPrime(x int) bool {
	mid := int(math.Sqrt(float64(x)))
	for i := 2; i <= mid+1; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
