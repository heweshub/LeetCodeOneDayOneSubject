package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	st := []string{}
	i := 0
	for i < len(s) {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			// 传指针就很妙
			digits := getDigits(s, &i)
			st = append(st, digits)
		} else if (ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z') || ch == '[' {
			st = append(st, string(ch))
			i++
		} else {
			i++
			sub := []string{}
			for st[len(st)-1] != "[" {
				sub = append(sub, st[len(st)-1])
				st = st[:len(st)-1]
			}
			for i := 0; i < len(st)/2; i++ {
				sub[i], sub[len(st)-i-1] = sub[len(st)-i-1], sub[i]
			}
			st = st[:len(st)-1]
			repeat, _ := strconv.Atoi(st[len(st)-1])
			st = st[:len(st)-1]
			t := strings.Repeat(getString(sub), repeat)
			st = append(st, t)
		}
	}
	return getString(st)
}

func getDigits(s string, i *int) string {
	ret := ""
	for ; s[*i] >= '0' && s[*i] <= '9'; *i++ {
		ret += string(s[*i])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
