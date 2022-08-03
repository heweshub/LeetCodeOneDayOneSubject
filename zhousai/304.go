package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	idxs := make([]int, n)
	for i := range idxs {
		fmt.Scan(&idxs[i])
	}
	//     fmt.Println(nums,idxs)
	m := map[int][]int{}
	for i, v := range idxs {
		m[v] = append(m[v], nums[i])
	}
	fmt.Println(nums, idxs, m)
	dp := make([]int, n)
	copy(dp, nums)
	for i := n - 1; i >= 0; i-- {

	}
	fmt.Println(dp)
	ans := math.MinInt64

	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
