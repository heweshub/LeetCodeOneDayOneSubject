package main

import (
	"strconv"
)

func convertTen2bin(n int) string {
	s := ""
	for ; n > 0; n /= 2 {
		s = strconv.Itoa(n%2) + s
	}
	return s
}
func hasAlternatingBits(n int) (ans bool) {
	s := convertTen2bin(n)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}

// func main() {
// 	flag := hasAlternatingBits(11)
// 	fmt.Println(flag)
// }
