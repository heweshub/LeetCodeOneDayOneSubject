package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "asd asdjhoufhwof asd w s"
	ss := strings.ReplaceAll(s, " ", "%20")
	fmt.Println(ss)
}
