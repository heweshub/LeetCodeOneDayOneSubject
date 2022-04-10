package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	alpha := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....",
		"..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.",
		"...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	cntMap := make(map[string]int, 26)
	n := len(words)
	for i := 0; i < n; i++ {
		s := ""
		for _, v := range words[i] {
			s += alpha[int(byte(v)-'a')]
		}
		cntMap[s]++
	}
	return len(cntMap)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))
}
