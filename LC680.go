package main

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	flag := false
	res1, res2 := true, true
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			if !flag {
				flag = true
				j--
			} else {
				res1 = false
				break
			}
		}
	}
	i, j = 0, len(s)-1
	flag = false
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			if !flag {
				flag = true
				i++
			} else {
				res2 = false
				break
			}
		}
	}
	if res1 || res2 {
		return true
	}
	return false
}

func main() {
	validPalindrome("abc")
}
