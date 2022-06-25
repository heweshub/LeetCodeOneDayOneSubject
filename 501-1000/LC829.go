package main

func consecutiveNumbersSum(n int) int {
	ans := 0
	n *= 2
	for k := 1; k*k < n; k++ {
		if n%k != 0 {
			continue
		}
		if (n/k-(k-1))%2 == 0 {
			ans++
		}
	}
	return ans
}

// func main() {
// 	fmt.Println(consecutiveNumbersSum(15))
// }
