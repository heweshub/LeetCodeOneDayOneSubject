package main

func findMin(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		pivot := (h + l) / 2
		if nums[pivot] < nums[h] {
			h = pivot
		} else {
			l = pivot + 1
		}
	}
	return nums[l]
}
