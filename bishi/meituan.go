package main

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	names := make([]string, n)
// 	for i := range names {
// 		fmt.Scan(&names[i])
// 	}
// 	for _, str := range names {
// 		flag := false
// 		for j, b := range str {
// 			isLetter := false
// 			isNumber := false
// 			if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' {
// 				isLetter = true
// 			}
// 			if b >= '0' && b <= '9' {
// 				isNumber = true
// 			}
// 			if j == 0 && !isLetter {
// 				break
// 			}
// 			if j > 0 && isNumber {
// 				flag = true
// 			}
// 			if j > 0 && !isLetter && !isNumber {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag {
// 			fmt.Println("Accept")
// 		} else {
// 			fmt.Println("Wrong")
// 		}
// 	}
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	nums := make([]int, n)
// 	for i := range nums {
// 		fmt.Scan(&nums[i])
// 	}
// 	indexs := make([]int, n)
// 	for i := range indexs {
// 		fmt.Scan(&indexs[i])
// 	}

// 	preSum := make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		preSum[i+1] = nums[i] + preSum[i]
// 	}

// }
