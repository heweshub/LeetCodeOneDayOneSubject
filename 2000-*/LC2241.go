package main

var a = [5]int{20, 50, 100, 200, 500}

type ATM [5]int

func Constructor() ATM { return ATM{} }

func (left *ATM) Deposit(bank []int) {
	for i, count := range bank {
		left[i] += count
	}
}

func (left *ATM) Withdraw(amount int) []int {
	ans := make([]int, 5)
	for i := 4; i >= 0; i-- {
		ans[i] = min(amount/a[i], left[i])
		amount -= ans[i] * a[i]
	}
	if amount > 0 {
		return []int{-1}
	}
	for i, count := range ans {
		left[i] -= count
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
