package main

func rotateString(s string, goal string) bool {
	fisrt := []int{}
	// 找到相对应的字符然后旋转就可以啦
	for i := 0; i < len(goal); i++ {
		if s[0] == goal[i] {
			fisrt = append(fisrt, i)
		}
	}
	for i := 0; i < len(fisrt); i++ {
		newGoal := goal[fisrt[i]:] + goal[:fisrt[i]]
		if newGoal == s {
			return true
		}
	}
	return false
}
