package main

func totalNQueens(n int) (ans int) {
	columns := make([]bool, n)
	dia1 := make([]bool, 2*n-1)
	dia2 := make([]bool, 2*n-1)
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}
		for col, hasQ := range columns {
			d1, d2 := row+n-1-col, row+col
			if hasQ || dia1[d1] || dia2[d2] {
				continue
			}
			// 回溯的关键
			columns[col] = true
			dia1[d1] = true
			dia2[d2] = true
			backtrack(row + 1)
			columns[col] = false
			dia1[d1] = false
			dia2[d2] = false
		}
	}
	backtrack(0)
	return
}

// func main() {
// 	fmt.Println(totalNQueens(5))
// }
