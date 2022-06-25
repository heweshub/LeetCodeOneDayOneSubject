package main

func duplicateZeros(arr []int) {
	i := 0
	for i < len(arr) {
		if arr[i] == 0 {
			if i+1 < len(arr) {
				j := len(arr) - 1
				for j > i {
					arr[j] = arr[j-1]
					j--
				}
			}
			i += 2
		} else {
			i++
		}
	}
}
