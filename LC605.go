package main

import "fmt"

func canPlaceFlowers(flowerbed []int, k int) bool {
	n := len(flowerbed)
	ans := 0
	isPlant := make([]bool, n)

	for i := 0; i < n; i++ {
		if flowerbed[i] == 1 {
			isPlant[i] = true
			if i-1 >= 0 {
				isPlant[i-1] = true
			}
			if i+1 < n {
				isPlant[i+1] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		if flowerbed[i] == 0 && isPlant[i] == false {
			ans++
			isPlant[i] = true
			if i-1 >= 0 {
				isPlant[i-1] = true
			}
			if i+1 < n {
				isPlant[i+1] = true
			}
		}
	}
	fmt.Println(isPlant)
	return ans == k
}

// func main() {
// 	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
// }
