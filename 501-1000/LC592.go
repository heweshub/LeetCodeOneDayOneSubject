package main

import (
	"fmt"
	"unicode"
)

func fractionAddtion(expression string) string {
	fenzi, fenmu := 0, 1
	for i, n := 0, len(expression); i < n; {
		fenzi1, sign := 0, 1
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}
		for i < n && unicode.IsDigit(rune(expression[i])) {
			fenzi1 = fenzi1*10 + int(expression[i]-'0')
			i++
		}
		fenzi1 *= sign
		// 跳过/
		i++

		fenmu1 := 0
		for i < n && unicode.IsDigit(rune(expression[i])) {
			fenmu1 = fenmu1*10 + int(expression[i]-'0')
			i++
		}
		fenzi = fenzi*fenmu1 + fenzi1*fenmu
		fenmu *= fenmu1
	}
	if fenzi == 0 {
		return "0/1"
	}
	g := gcd(abs(fenzi), fenmu)
	return fmt.Sprintf("%d/%d", fenzi/g, fenmu/g)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
