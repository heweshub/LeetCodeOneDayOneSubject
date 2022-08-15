package main

func canFinish(n int, nums [][]int) bool {
	var (
		edges  = make([][]int, n)
		vis    = make([]int, n)
		result []int
		valid  = true
		dfs    func(u int)
	)
	dfs = func(u int) {
		vis[u] = 1
		for _, v := range edges[u] {
			if vis[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if vis[v] == 1 {
				valid = false
				return
			}
		}
		vis[u] = 2
		result = append(result, u)
	}
	for _, info := range nums {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < n && valid; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
