package main

// func leastInterval(t []byte, n int) (minTime int) {
// 	// 统计相同任务的个数
// 	cnt := map[byte]int{}
// 	for _, item := range t {
// 		cnt[item]++
// 	}
// 	// 初始化下一个有效位置和剩余任务数
// 	nextValid := make([]int, 0, len(cnt))
// 	rest := make([]int, 0, len(cnt))
// 	for _, c := range cnt {
// 		nextValid = append(nextValid, 1)
// 		rest = append(rest, c)
// 	}
// 	//
// 	for range t {
// 		minTime++
// 		minNextValid := math.MaxInt64
// 		for i, r := range rest {
// 			if r > 0 && nextValid[i] < minNextValid {
// 				minNextValid = nextValid[i]
// 			}
// 		}
// 		if minNextValid > minTime {
// 			minTime = minNextValid
// 		}
// 		best := -1
// 		for i, r := range rest {
// 			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
// 				best = i
// 			}
// 		}
// 		nextValid[best] = minTime + n + 1
// 		rest[best]--
// 	}
// 	return
// }

func leastInterval(t []byte, n int) int {
	arr := make([]int, 26)
	for _, v := range t {
		arr[v-'A']++
	}
	max := maxOfArray(arr)
	ret := (max - 1) * (n + 1)
	for i := 0; i < 26; i++ {
		if arr[i] == max {
			ret++
		}
	}
	if ret > len(t) {
		return ret
	}
	return len(t)
}

// func maxOfArray(a []int) (ans int) {
// 	for _, v := range a {
// 		if v > ans {
// 			ans = v
// 		}
// 	}
// 	return
// }
