package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	newWords := make([]string, len(words))
	for i := range words {
		var newWord string
		if strings.ToLower(string(words[i][0])) == "a" ||
			strings.ToLower(string(words[i][0])) == "e" ||
			strings.ToLower(string(words[i][0])) == "i" ||
			strings.ToLower(string(words[i][0])) == "o" ||
			strings.ToLower(string(words[i][0])) == "u" {
			newWord = words[i] + "ma"
		} else {
			newWord = words[i][1:] + string(words[i][0]) + "ma"
		}
		for j := 1; j <= i+1; j++ {
			newWord += "a"
		}
		newWords[i] = newWord
	}
	return strings.Join(newWords, " ")
}

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}
