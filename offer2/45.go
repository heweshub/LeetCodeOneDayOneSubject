package main

import (
	"strconv"
)

func minNumber(nums []int) string {
	res := make([]string, len(nums))
	for i, v := range nums {
		res[i] = strconv.Itoa(v)
	}
	// 比较组成的数字的大小
	compare := func(str1, str2 string) bool {
		num1, _ := strconv.Atoi(str1 + str2)
		num2, _ := strconv.Atoi(str2 + str1)
		if num1 < num2 {
			return true
		}
		return false
	}

	var quickSort func([]string, int, int)
	quickSort = func(str []string, left, right int) {
		if left >= right {
			return
		}
		l, r, pivot := left, right, str[left]
		for l < r {
			for l < r && compare(pivot, str[r]) {
				r--
			}
			for l < r && compare(str[l], pivot) {
				l++
			}
			str[l], str[r] = str[r], str[l]
		}
		str[l], str[left] = str[left], str[l]
		quickSort(str, left, l-1)
		quickSort(str, l+1, right)
	}
	quickSort(res, 0, len(res)-1)
	ans := ""
	for _, s := range res {
		ans += s
	}
	return ans
}
