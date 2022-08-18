package main

import "strings"

func removeKdigits(num string, k int) string {
	st := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(st) > 0 && digit < st[len(st)-1] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, digit)
	}
	st = st[:len(st)-1]
	ans := strings.TrimLeft(string(st), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
