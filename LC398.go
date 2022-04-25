package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	mmap map[int][]int
}

func Constructor2(nums []int) Solution {
	m := map[int][]int{}
	for i := range nums {
		m[nums[i]] = append(m[nums[i]], i)
	}
	return Solution{mmap: m}
}

func (s *Solution) Pick(target int) int {
	return s.mmap[target][rand.Intn(len(s.mmap[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

func main() {
	nums := []int{1, 2, 3, 3, 3}
	obj := Constructor2(nums)
	param1 := obj.Pick(2)
	fmt.Println(param1)
}
