package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	cnt := map[string]int{}
	for _, s := range cpdomains {
		i := strings.IndexByte(s, ' ')
		c, _ := strconv.Atoi(s[:i])
		s = s[i+1:]
		cnt[s] += c
		for {
			i := strings.IndexByte(s, '.')
			if i < 0 {
				break
			}
			s = s[i+1:]
			cnt[s] += c
		}
	}
	ans := make([]string, 0, len(cnt))
	for s, c := range cnt {
		ans = append(ans, strconv.Itoa(c)+" "+s)
	}
	return ans
}
