package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "sss,sss,sss"
	a := strings.Split(s, ",")
	fmt.Println(a)
}
