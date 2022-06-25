package main

import (
	"unicode"
)

type NestedInteger struct {
}

func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		// num, _ := strconv.Atoi(s)
		ni := &NestedInteger{}
		// ni.SetInteger(num)
		return ni
	}
	stack, num, negative := []*NestedInteger{}, 0, false
	for i, ch := range s {
		if ch == '-' {
			negative = true
		} else if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		} else if ch == '[' {
			stack = append(stack, &NestedInteger{})
		} else if ch == ',' || ch == ']' {
			if unicode.IsDigit(rune(s[i-1])) {
				if negative {
					num = -num
				}
				// ni := NestedInteger{}
				// ni.SetInteger(num)
				// stack[len(stack)-1].Add(ni)
			}
			num, negative = 0, false
			if ch == ']' && len(stack) > 1 {
				// stack[len(stack)-2].Add(*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack[len(stack)-1]
}
