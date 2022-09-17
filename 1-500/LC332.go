package main

import "sort"

func findItinerary(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)
	for _, t := range tickets {
		src, dst := t[0], t[1]
		m[src] = append(m[src], dst)
	}

	for key := range m {
		sort.Strings(m[key])
	}
	var dfs func(cur string)
	dfs = func(cur string) {
		for {
			if v, ok := m[cur]; !ok || len(v) == 0 {
				break
			}
			tmp := m[cur][0]
			m[cur] = m[cur][1:]
			dfs(tmp)
		}
		res = append(res, cur)
	}
	dfs("JFK")
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}
