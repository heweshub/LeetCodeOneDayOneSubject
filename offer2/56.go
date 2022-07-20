package main

func singleNumbers(nums []int) []int {
	var a int
	for _, v := range nums {
		a ^= v
	}
	mask := a & (-a)
	ans := make([]int, 2)
	for _, v := range nums {
		if v&mask == 0 {
			ans[0] ^= v
		} else {
			ans[1] ^= v
		}
	}
	return ans
}
