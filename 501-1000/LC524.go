package main

import "sort"

func findLongestWord(s string, dictionary []string) (ans string) {
	max := 0
	sort.Slice(dictionary, func(i, j int) bool {
		return dictionary[i] < dictionary[j]
	})
	for _, word := range dictionary {
		i, j := 0, 0
		for i < len(s) && j < len(word) {
			if s[i] == word[j] {
				i++
				j++
			} else {
				i++
			}
			if j == len(word) {
				if max < j {
					max = j
					ans = word
				}
				break
			}
		}
	}
	return
}
