package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	words := strings.Fields(s)
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		if i != 0 {
			res += " "
		}
	}
	return res
}
