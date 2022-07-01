package main

func heapSort(nums []int) []int {
	var heapify func([]int, int, int)
	heapify = func(nums []int, root, end int) {
		for {
			child := root*2 + 1
			if child > end {
				return
			}
			if child < end && nums[child] <= nums[child+1] {
				child++
			}
			if nums[root] > nums[child] {
				return
			}
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		}
	}
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify(nums, 0, end)
	}
	return nums
}
