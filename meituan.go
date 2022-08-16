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
