package main

type MyCalendarTwo struct {
	calendar, overlaps [][]int
}

// func Constructor() MyCalendarTwo {
// 	return MyCalendarTwo{[][]int{}, [][]int{}}
// }

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, v := range this.overlaps {
		if start < v[1] && end > v[0] {
			return false
		}
	}
	for _, v := range this.calendar {
		if start < v[1] && end > v[0] {
			this.overlaps = append(this.overlaps, []int{max(start, v[0]), min(end, v[1])})
		}
	}
	this.calendar = append(this.calendar, []int{start, end})
	return true
}

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
