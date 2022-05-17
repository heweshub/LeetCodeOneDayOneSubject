package main

import (
	"fmt"
	"strings"
)

func isAlienSorted(words []string, order string) bool {
	n := len(words)
	for i := 0; i < n-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		for j := 0; j < len(word1); j++ {
			if j >= len(word2) {
				return false
			}
			id1 := strings.IndexByte(order, word1[j])
			id2 := strings.IndexByte(order, word2[j])
			if id1 < id2 {
				break
			} else if id1 > id2 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
}
