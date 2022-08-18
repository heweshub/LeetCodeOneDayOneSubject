package main

func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, v := range nums {
		maxVal = max(maxVal, v)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	return rob(sum)
}
func rob(nums []int) int {
	f, s := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f, s = s, max(f+nums[i], s)
	}
	return s
}

// func max(a ...int) int {
// 	ans := a[0]
// 	for _, v := range a[1:] {
// 		if ans < v {
// 			ans = v
// 		}
// 	}
// 	return ans
// }
