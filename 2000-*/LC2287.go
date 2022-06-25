package main

import "math"

func rearrangeCharacters(s string, target string) int {
	m := make(map[byte]int)
	mm := make(map[byte]int)
	for _, ch := range s {
		m[byte(ch)]++
	}
	for _, ch := range target {
		mm[byte(ch)]++
	}
	min := math.MaxInt32
	for k, v := range mm {
		if m[k]/v < min {
			min = m[k] / v
		}
	}
	return min
}
