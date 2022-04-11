package main

import (
	"fmt"
	"time"
)

var cnts = []int{10, 101, 840, 6115, 38606, 207177, 920068}

//func countNumbersWithUniqueDigits(n int) int {
//	if n == 1 {
//		return 9
//	}
//	ncr := math.Pow(10, float64(n))
//	ncl := math.Pow(10, float64(n-1))
//
//	cnt := 0
//	for i := int(ncl); i < int(ncr); i++ {
//		mmap := make(map[string]int, 10)
//		s := strconv.Itoa(i)
//		ss := strings.Split(s, "")
//		for j := 0; j < len(ss); j++ {
//			mmap[ss[j]]++
//		}
//		if len(ss) == len(mmap) {
//			cnt++
//		}
//	}
//	return cnt+cnts[n-2]
//}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	ans, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		ans += cur
	}
	return ans
}
func main() {
	fmt.Println(time.Now())
	fmt.Println(countNumbersWithUniqueDigits(8))
	fmt.Println(time.Now())
}
