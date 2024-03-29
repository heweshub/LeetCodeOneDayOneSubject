package main

func largestRectangleArea(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && height[mono_stack[len(mono_stack)-1]] >= height[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*height[i])
	}
	return ans
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
