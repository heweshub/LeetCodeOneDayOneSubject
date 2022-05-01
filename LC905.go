package main

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	newNums := make([]int, n)
	i, j := 0, n-1
	for k := 0; k < n; k++ {
		if nums[k]&1 == 1 {
			newNums[j] = nums[k]
			j--
		} else {
			newNums[i] = nums[k]
			i++
		}
	}
	return newNums
}

func main() {
	sortArrayByParity([]int{3, 1, 2, 4})
}
