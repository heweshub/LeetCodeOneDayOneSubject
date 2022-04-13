package main

import (
	"math/rand"
)

type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor1() RandomizedSet {
	return RandomizedSet{
		[]int{}, map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.indices[val]
	if !ok {
		return false
	}
	// 将要删除的下标添上nums最后的元素，相应的indices也要更新，最后修改nums和indices
	last := len(this.nums) - 1
	this.nums[id] = this.nums[last]
	this.indices[this.nums[id]] = id
	this.nums = this.nums[:last]
	delete(this.indices, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

//["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
//[[], [1], [2], [2], [], [1], [2], []]
//输出
//[null, true, false, true, 2, true, false, 2]
