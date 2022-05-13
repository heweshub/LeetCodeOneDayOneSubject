package main

import "fmt"

func minDeletionSize(strs []string) (cnt int) {
	n, m := len(strs), len(strs[0])
	mm := make([][]string, m)
	for i := range mm {
		mm[i] = make([]string, n)
	}
	for i, v1 := range strs {
		for j, v2 := range v1 {
			mm[j][i] = string(v2)
		}
	}
	for i := range mm {
		for j := 0; j < n-1; j++ {
			if mm[i][j] > mm[i][j+1] {
				cnt++
				break
			}
		}
	}
	return
}
func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	fmt.Println(minDeletionSize([]string{"a", "b"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
