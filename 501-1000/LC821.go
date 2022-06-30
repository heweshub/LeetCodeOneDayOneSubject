package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		flagl := false
		flagr := false
		if s[i] == 'e' {
			ans[i] = 0
		} else {
			j := i
			for j < len(s)-1 {
				j++
				if s[j] == 'e' {
					flagr = true
					break
				}
			}
			k := i
			for k > 0 {
				k--
				if s[k] == 'e' {
					flagl = true
					break
				}
			}
			if flagl && flagr {
				ans[i] = min1(j-i, i-k)
			} else if flagl {
				ans[i] = k - i
			} else {
				ans[i] = j - i
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func min1(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// func main() {
// 	shortestToChar("loveleetcode", 'e')
// 	// [4 3 2 0 0 0 0 0 -1 -2 -3 0]
// 	// [3,2,1,0,1,0,0,1, 2, 2, 1,0]
// }
