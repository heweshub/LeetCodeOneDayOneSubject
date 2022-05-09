package main

func findDuplicates(nums []int) (ans []int) {
	// nums[i] 放在对应下标-1的位置
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	// 只出现一次的元素已经放置在对应位置，下标与对应的值不相等即为重复元素
	for i, num := range nums {
		if num-1 != i {
			ans = append(ans, num)
		}
	}
	return
}

func main() {
	findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
}
