package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int, k int) int {
	n := len(nums)
	for i := 0; i < k; i++ {
		sort.Ints(nums)
		nums[0]++
	}
	fmt.Println(nums)
	multi := 1
	for i := 0; i < n; i++ {
		multi *= nums[i]
	}
	return multi % (1e9 + 7)
}

//[92,36,15,84,57,60,72,86,70,43,16]
//62
func main() {
	fmt.Println(maximumProduct([]int{92, 36, 15, 84, 57, 60, 72, 86, 70, 43, 16}, 62))
}
