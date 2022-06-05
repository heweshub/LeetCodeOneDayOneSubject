package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) (ans int) {
	n := len(emails)
	m := make(map[string]bool, n)
	for i := range emails {
		email := strings.Split(emails[i], "@")
		copyemail := ""
		for _, v := range email[0] {
			if v == '.' {
				continue
			}
			if v == '+' {
				break
			}
			copyemail += string(v)
		}
		if _, exist := m[copyemail+"@"+email[1]]; !exist {
			ans++
			m[copyemail+"@"+email[1]] = true
		}
	}
	return
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
