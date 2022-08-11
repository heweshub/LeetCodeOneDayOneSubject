package main

import "fmt"

func quickSort(nums []int) []int {
	var quick func(nums []int, left, right int) []int
	quick = func(nums []int, left, right int) []int {
		l, r := left, right
		if l > r {
			return nil
		}
		pivot := nums[l]
		for l < r {
			for l < r && nums[r] >= pivot {
				r--
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
			for l < r && nums[l] <= pivot {
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

func main() {
	fmt.Println(quickSort([]int{7, 3, 8, 2, 5, 1, 9}))
}
