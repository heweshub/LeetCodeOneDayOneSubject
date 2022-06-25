package main

func selfDividingNumbers(left int, right int) []int {
	if left > right {
		return []int{}
	}
	var cnt []int
	for i := left; i <= right; i++ {
		if check(i) {
			cnt = append(cnt, i)
		}
	}
	return cnt
}

func check(num int) bool {
	n := num
	for n > 0 {
		yuShu := n % 10
		if yuShu == 0 || num%yuShu != 0 {
			return false
		}
		n /= 10
	}
	return true
}

// func main() {
// 	fmt.Println(selfDividingNumbers(47, 85))
// }
