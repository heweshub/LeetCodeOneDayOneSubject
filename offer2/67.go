package main

import (
	"strings"
	"unicode"
)

const INT_MAX = 2147483648

func strToInt(str string) (ans int) {

	s := strings.Trim(str, " ")
	if len(s) == 0 {
		return 0
	}
	flag := false
	i := 0
	if !unicode.IsDigit(rune(s[i])) && s[i] != '+' && s[i] != '-' {
		return 0
	}
	if s[i] == '-' {
		flag = true
		i++
	} else if s[i] == '+' {
		i++
	}
	for i < len(s) {
		if !unicode.IsDigit(rune(s[i])) {
			break
		}
		ans = ans*10 + int(s[i]-'0')
		if ans > INT_MAX {
			ans = INT_MAX
			break
		}
		i++
	}
	if flag {
		ans *= -1
	}
	if ans == INT_MAX {
		ans = INT_MAX - 1
	}
	return
}
