package main

import "fmt"

type pair struct {
	v, t int
}

func totalSteps(nums []int) (ans int) {
	stack := []pair{}
	for _, num := range nums {
		step := 0
		for len(stack) > 0 && stack[len(stack)-1].v <= num {
			step = max(step, stack[len(stack)-1].t)
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			step++
			ans = max(ans, step)
		} else {
			step = 0
		}
		fmt.Print(step)
		stack = append(stack, pair{num, step})
	}
	return
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// func main() {
// 	totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11})
// }
