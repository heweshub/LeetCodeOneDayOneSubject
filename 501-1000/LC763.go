package main

func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	// 分别标记出哥各个字母最后出现的下标位置
	for i, c := range s {
		lastPos[c-'a'] = i
	}

	start, end := 0, 0
	for i, c := range s {
		// 找到最大的end然后分割
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}

// func main() {
// 	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
// }
