package main

// N皇后问题
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	dias1, dias2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, dias1, dias2)
	return solutions
}

func backtrack(queens []int, n, row int, columns, dias1, dias2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		dia1 := row - i
		if dias1[dia1] {
			continue
		}
		dia2 := row + i
		if dias2[dia2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		dias1[dia1], dias2[dia2] = true, true
		backtrack(queens, n, row+1, columns, dias1, dias2)
		queens[row] = -1
		delete(columns, i)
		delete(dias1, dia1)
		delete(dias2, dia2)
	}
}

func generateBoard(queens []int, n int) (board []string) {
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return
}

// func main() {
// 	for _, i := range solveNQueens(8) {
// 		for _, v := range i {
// 			fmt.Println(v)
// 		}
// 	}
// }
