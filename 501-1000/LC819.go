package main

import (
	"fmt"
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	words := selectLetter(paragraph)
	n := len(words)
	mmap := make(map[string]int, n)
	for _, v := range words {
		v = strings.ToLower(v)
		mmap[v]++
	}
	for i := 0; i < len(banned); i++ {
		delete(mmap, banned[i])
	}
	max := 0
	ans := ""
	for k, v := range mmap {
		if v > max {
			max = v
			ans = k
		}
	}
	fmt.Println(ans)
	return ans
}

func selectLetter(req string) (ans []string) {
	n := len(req)
	i := 0
	begin := i
	for i < n {
		if i == n-1 && unicode.IsLetter(rune(req[i])) {
			ans = append(ans, req[begin:i+1])
		}
		if !unicode.IsLetter(rune(req[i])) {
			reqStr := req[begin:i]
			if reqStr == "" {
				begin = i + 1
				i++
				continue
			}
			ans = append(ans, reqStr)
			begin = i + 1
		}
		i++
	}
	return
}

// func main() {
// 	mostCommonWord("Bob!", []string{"hit"})
// }
