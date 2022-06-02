package main

func sortArray(nums []int) []int {
	var quick func([]int, int, int) []int
	quick = func(nums []int, left, right int) []int {
		if left > right {
			return nil
		}
		l, r := left, right
		pivot := nums[left]
		for l < r {
			for l < r && pivot <= nums[r] {
				r--
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
			for l < r && pivot > nums[l] {
				l++
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		nums[l] = pivot
		quick(nums, left, l-1)
		quick(nums, l+1, right)
		return nums
	}
	return quick(nums, 0, len(nums)-1)
}
