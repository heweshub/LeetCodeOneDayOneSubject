package main

import "fmt"

func maopaosort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] >= nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
}

func main() {
	//nums := []int{6, 3, 9, 2, 7, 1, 8, 4}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2}
	maopaosort(nums)
}
