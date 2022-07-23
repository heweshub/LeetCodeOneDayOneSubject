package main

import (
	"fmt"
	"math/big"
	"strings"
)

func isNumber(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	if strings.ContainsAny(s, " bppx") {
		return false
	}
	if i := strings.IndexAny(s, "e"); i >= 0 {
		n1, _ := fmt.Sscan(s[:i], new(big.Float), new(string))
		n2, _ := fmt.Sscanf(s[i+1:], "%d%s", new(big.Int), new(string))
		return n1 == 1 && n2 == 1
	}
	n, _ := fmt.Sscan(s, new(big.Float), new(string))
	return n == 1
}
