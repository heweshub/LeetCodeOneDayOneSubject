package main

func distributeCookies(cookies []int, k int) int {
	ret := 1<<31 - 1

	records := make([]int, k)

	var dfs func(int)
	dfs = func(i int) {
		if i == len(cookies) {
			// 全部分配完cookies后看records中的最大值，记录每一次最大值。
			curMax := 0
			for _, v := range records {
				if v > curMax {
					curMax = v
				}
			}
			if curMax < ret {
				ret = curMax
			}
			return
		}
		// 回溯算法
		for j := range records {
			records[j] += cookies[i]
			if records[j] < ret {
				dfs(i + 1)
			}
			records[j] -= cookies[i]
		}
	}
	dfs(0)
	return ret
}
