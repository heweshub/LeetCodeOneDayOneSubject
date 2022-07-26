package main

func isVaild(s string) bool {
	cnt := 0
	for _, ch := range s {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func helper(ans *[]string, s string, start, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isVaild(s) {
			*ans = append(*ans, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] {
			continue
		}
		if lremove+rremove > len(s)-i {
			return
		}
		if lremove > 0 && s[i] == '(' {
			helper(ans, s[:i]+s[i+1:], i, lremove-1, rremove)
		}
		if rremove > 0 && s[i] == ')' {
			helper(ans, s[:i]+s[i+1:], i, lremove, rremove-1)
		}
	}
}

func removeInvalidParemtheses(s string) (ans []string) {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}
	helper(&ans, s, 0, lremove, rremove)
	return
}
