package main

import "fmt"

// 全排列问题
var res [][]int
var depth = 0

func permute(nums []int) [][]int {
	back(nums, depth, len(nums))
	return res
}

func back(nums []int, dep, length int) {
	if dep == length {
		tmp := make([]int, length)
		copy(tmp, nums)
		res = append(res, tmp)
	}
	for i := dep; i < length; i++ {
		nums[dep], nums[i] = nums[i], nums[dep]
		back(nums, dep+1, length)
		nums[i], nums[dep] = nums[dep], nums[i]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}
