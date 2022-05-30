package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if sp := strings.Split(queryIP, "."); len(sp) == 4 {
		for _, s := range sp {
			if len(s) > 1 && s[0] == '0' {
				return "Neither"
			}
			if v, err := strconv.Atoi(s); err != nil || v > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if sp := strings.Split(queryIP, ":"); len(sp) == 8 {
		for _, s := range sp {
			if len(s) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseInt(s, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
}
