package main

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		mul := 1
		for j := i; j < n; j++ {
			if mul*nums[j] < k {
				ans++
				mul *= nums[j]
			} else {
				break
			}
		}
	}
	return
}

func numSubarrayProductLessThanK1(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return
}
