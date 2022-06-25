package main

func numberOfLines(widths []int, s string) (ans []int) {
	n := len(s)
	cnt := 0
	lines := 1
	for i := 0; i < n; i++ {
		cnt += widths[s[i]-'a']
		if cnt > 100 {
			cnt = widths[s[i]-'a']
			lines++
		}
	}
	//fmt.Println(lines, cnt)
	return []int{lines, cnt}
}

// func main() {
// 	numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa")
// }
