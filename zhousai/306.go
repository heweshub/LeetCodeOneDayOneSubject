package main

import "strconv"

func smallestNumber(pattern string) (ans string) {
	vis := make([]bool, 10)
	vis[0] = true
	var dfs func(int, string)
	dfs = func(index int, path string) {
		if len(path) == len(pattern)+1 {
			ans = path
			return
		}
		for i := 1; i < 10; i++ {
			if len(path) == 0 {
				path += string(i)
			}
			num, _ := strconv.Atoi(path[:len(path)-1])
			ch := path[index]
			if ch == 'I' {
				if len(path) > 0 && i < num {
					if !vis[i] {
						dfs(index+1, path+string(i))
						vis[i] = false
					}
				}
			} else {
				if len(path) > 0 && i > num {
					if !vis[i] {
						dfs(index+1, path+string(i))
						vis[i] = false
					}
				}
			}
			path = path[:len(path)-1]
		}
	}
	dfs(0, "")
	return
}
