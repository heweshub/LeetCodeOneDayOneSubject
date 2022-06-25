package main

func addStrings(num1 string, num2 string) (ans string) {
	n1, n2 := len(num1), len(num2)
	for n1 > 0 || n2 > 0 {
		x, y := 0, 0
		if n1 > 0 {
			x = int(num1[n1-1] - '0')
			n1--
		}
		if n2 > 0 {
			y = int(num2[n2-1] - '0')
			n2--
		}
		ans = string(rune(x+y+'0')) + ans
	}
	return
}

// func main() {
// 	fmt.Println("ans:", addStrings("123", "1"))
// }
