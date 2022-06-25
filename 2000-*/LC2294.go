package main

import "sort"

func partitionArray(nums []int, k int) (ans int) {
	sort.Ints(nums)
	ans, min := 1, nums[0]
	for _, num := range nums {
		if num-min > k {
			ans++
			min = num
		}
	}
	return
}

func main() {
	//partitionArray([]int{3, 6, 1, 2, 5}, 2)
	partitionArray([]int{1, 2, 3}, 1)
}
