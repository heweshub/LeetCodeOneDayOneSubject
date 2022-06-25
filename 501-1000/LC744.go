package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, v := range letters {
		if v <= target {
			continue
		} else {
			fmt.Println(byte(v))
			return byte(v)
		}
	}
	return byte(' ')
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}
