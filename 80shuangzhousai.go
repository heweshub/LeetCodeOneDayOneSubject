package main

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) (ans []int) {
	sort.Ints(potions)
	for _, v := range spells {
		var findNum int64
		if success%int64(v) == 0 {
			findNum = success / int64(v)
		} else {
			findNum = success/int64(v) + 1
		}
		cnt := erfen(potions, int(findNum))
		ans = append(ans, len(potions)-cnt)
	}
	return
}

func erfen(nums []int, t int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == t {
			return mid
		} else if nums[mid] > t {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

// func main() {
// 	//spells = [3,1,2], potions = [8,5,8], success = 16
// 	//fmt.Println(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
// 	//spells = [5,1,3], potions = [1,2,3,4,5], success = 7
// 	//fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
// 	fmt.Println('z' + 1)
// }
