package main

func buy(P, N, D, X, Y int) int {
	daynum := 0
	if D/X < N/Y {
		daynum = D / X
	} else {
		daycost := X + Y*P
		totalMoney := D + N*P
		daynum = totalMoney / daycost
	}
	return daynum
}
