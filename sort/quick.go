package main

func quickSort(nums []int) []int {
	var quick func(nums []int, left, right int) []int
	quick = func(nums []int, left, right int) []int {
		l, r := left, right
		if left > right {
			return nil
		}
		pivot := nums[left]
		for l < r {
			for l < r && nums[r] >= pivot {
				r--
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
			for l < r && nums[r] <= pivot {
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
