package main

func sortArray(nums []int) []int {
	merge := func(left, right []int) []int {
		res := make([]int, len(left)+len(right))
		var l, r, i int
		for l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				res[i] = left[l]
				l++
			} else {
				res[i] = right[r]
				r++
			}
			i++
		}
		copy(res[i:], left[l:])
		copy(res[i:], right[r:])
		return res
	}
	var sort func(nums []int) []int
	sort = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}
		mid := len(nums) / 2
		left := sort(nums[:mid])
		right := sort(nums[mid:])
		return merge(left, right)
	}
	return sort(nums)
}
