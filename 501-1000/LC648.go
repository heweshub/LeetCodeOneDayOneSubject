package main

import (
	"fmt"
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})
	fmt.Println(dictionary)
	m := make(map[string]bool, len(dictionary))
	for _, v := range dictionary {
		m[v] = true
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if _, ok := m[string(word[:j])]; ok {
				words[i] = word[:j]
				break
			}
		}

	}
	return strings.Join(words, " ")
}

func replaceWords1(dictionary []string, sentence string) string {
	type trie map[rune]trie // 定义前缀树
	root := trie{}

	for _, s := range dictionary {
		cur := root
		for _, c := range s {
			if cur[c] == nil {
				cur[c] = trie{}
			}
			cur = cur[c]
		}
		cur['#'] = trie{}
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, c := range word {
			if cur['#'] != nil {
				words[i] = word[:j]
				break
			}
			if cur[c] == nil {
				break
			}
			cur = cur[c]
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"catt", "cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
