package main

import (
	"strconv"
)

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit++
		count = 9 * start * digit
	}
	num := start + (n-1)/digit
	// fmt.Println(num, string(num))
	s := strconv.Itoa(num)
	return int(s[(n-1)%digit])
	// return num
}

// func main() {
// 	findNthDigit(11)
// 	// fmt.Println(findNthDigit(3))
// }
