package main

func validateStackSequences(pushed, poped []int) bool {
	st := []int{}
	j := 0
	for _, v := range pushed {
		st = append(st, v)
		for len(st) > 0 && st[len(st)-1] == poped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}
