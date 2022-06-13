package main

import "fmt"

func calculateTax(brackets [][]int, income int) (res float64) {
	for i, v := range brackets {
		if i == 0 {
			if v[0] >= income {
				res = float64(income*v[1]) / 100.00
				break
			} else {
				res += float64(v[1] * v[0])
			}
		} else if income > v[0] {
			res += float64((v[0]-brackets[i-1][0])*v[1]) / 100.00
		} else {
			res += float64((income-brackets[i-1][0])*v[1]) / 100.00
			break
		}
	}
	fmt.Sprintf("%.5f", res)
	return
}
