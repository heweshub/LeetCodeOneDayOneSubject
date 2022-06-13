package main

import "fmt"

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	mp := ['z' + 1]['z' + 1]bool{}
	for _, p := range mappings {
		mp[p[0]][p[1]] = true
	}
next:
	for i := len(sub); i <= len(s); i++ {
		for j, c := range s[i-len(sub) : i] {
			if byte(c) != sub[j] && !mp[sub[j]][c] {
				continue next
			}
		}
		return true
	}
	return false
}

func main() {
	//s = "fool3e7bar", sub = "leet", mappings = [["e","3"],["t","7"],["t","8"]]
	fmt.Println(matchReplacement("fool3e7bar", "leet", [][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}}))
}
