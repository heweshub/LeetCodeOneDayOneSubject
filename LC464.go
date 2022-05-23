package main

func canIWin(maxChoosableInteger, desiredTotal int) bool {
	// 不能达到desiredTotal
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	dp := make([]int8, 1<<maxChoosableInteger)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(int, int) int8
	dfs = func(useNum, curTot int) (res int8) {
		dv := &dp[useNum]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := 0; i < maxChoosableInteger; i++ {
			if useNum>>i&1 == 0 && (curTot+i+1 >= desiredTotal || dfs(useNum|1<<i, curTot+i+1) == 0) {
				return 1
			}
		}
		return
	}
	return dfs(0, 0) == 1
}
