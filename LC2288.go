package main

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	newWords := []string{}
	for _, word := range words {
		if byte(word[0]) == '$' {
			if num, err := strconv.Atoi(word[1:]); err == nil {
				newNum := fmt.Sprintf("%.2f", float64(num)*(1.0-float64(discount)/100.0))
				word = "$" + string(newNum)
			}
		}
		newWords = append(newWords, word)
	}
	return strings.Join(newWords, " ")
}

func main() {
	discountPrices("there are $1 $2 and 5$ candies in the shop", 50)
}
