package main

import "fmt"

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	// map用于记录已经遍历过的基因序列
	// 直接遍历bank中的基因序列，其余未出现的均为无效的基因
	bankSet := map[string]struct{}{}
	for _, s := range bank {
		bankSet[s] = struct{}{}
	}
	// bank中不存在end，直接返回-1
	if _, ok := bankSet[end]; !ok {
		return -1
	}
	q := []string{start}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x {
						nxt := cur[:i] + string(y) + cur[i+1:]
						// 检测合法性
						if _, ok := bankSet[nxt]; ok {
							if nxt == end {
								return step + 1
							}
							delete(bankSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))

}
