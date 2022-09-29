package main

import (
	"container/heap"
	"sort"
)

var factors = []int{3, 5, 7}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func getKthMagicNumber(k int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}

	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == k {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}
