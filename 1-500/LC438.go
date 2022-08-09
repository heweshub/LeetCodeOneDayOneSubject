package main

func findAnagrams(s, p string) (ans []int) {
	ns, np := len(s), len(p)
	if ns < np {
		return
	}
	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	// 滑动窗口大小np
	for i, ch := range s[:ns-np] {
		sCount[ch-'a']--
		sCount[s[i+np]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
