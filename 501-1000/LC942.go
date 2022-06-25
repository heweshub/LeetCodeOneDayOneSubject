package main

func diStringMatch(s string) (ans []int) {
	l, r := 0, len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			ans = append(ans, l)
			l++
		} else {
			ans = append(ans, r)
			r--
		}
	}
	ans = append(ans, l)
	return
}
