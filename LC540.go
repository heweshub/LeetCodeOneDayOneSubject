package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		mid := (h + l) / 2
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return nums[l]
}

func main() {
	a := 11
	fmt.Println(a ^ 1)
	fmt.Println(a + 1)
}
