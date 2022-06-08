package main

import "strings"

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	negative := false
	var abs string
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		abs = s
	case '-':
		negative = true
		abs = s[1:]
	case '+':
		abs = s[1:]
	default:
		return 0
	}
	absNum := 0
	for _, ch := range abs {

		if int(ch-'0') > 9 || int(ch-'0') < 0 {
			break
		}
		absNum = absNum*10 + int(ch-'0')
		if absNum >= 1<<31 {
			absNum = 1 << 31
			break
		}
	}
	if negative {
		absNum = -absNum
	} else {
		if absNum >= 1<<31-1 {
			absNum = 1<<31 - 1
		}
	}
	return absNum
}
