package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans := ""
	dfs("?12?0?9??", []byte{}, 0, &ans)
	fmt.Println(ans)
}

func dfs(s string, sb []byte, idx int, ans *string) {
	if idx >= len(s) {
		num, _ := strconv.Atoi(string(sb))
		if num%3 == 0 {
			*ans = string(sb)
		}
		return
	}
	if s[idx] != '?' {
		sb = append(sb, s[idx])
		dfs(s, sb, idx+1, ans)
		sb = sb[:len(sb)-1]
	} else {
		for i := 0; i < 10 && *ans == ""; i++ {
			if i == 0 && idx == 0 {
				continue
			}
			if idx-1 >= 0 && int(sb[idx-1]-'0') == i {
				continue
			}
			if idx+1 < len(s) && s[idx+1] != '?' && int(s[idx+1]-'0') == i {
				continue
			}
			sb = append(sb, byte(i+48))
			dfs(s, sb, idx+1, ans)
			sb = sb[:len(sb)-1]
		}
	}
}
