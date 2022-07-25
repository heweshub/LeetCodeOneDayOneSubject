package main

import "fmt"


var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	dia1, dia2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, dia1, dia2)
	return solutions
}

func backtrack(queens []int, n,row int, columns,dia1,dia2, map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i:=0; i<n; i++ {
		// 列不重复
		if columns[i] {
			continue
		}
		// 右斜线上不重复
		di1 := row-i
		if dia1[di1] {
			continue
		}
		// 左斜线上不重复
		di2 := row + i
		if dia2[di2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		dia1[di1],dia2[di2] = true, true
		backtrack(queens,n,row+1,columns,dia1,dia2)
		queens[row] = -1
		delete(columns,i)
		delete(dia1,di1)
		delete(dia2,di2)
	}
}

func generateBoard(queens []int, n int) []string{
	board := []string{}
	for i:=0;i<n;i++{
		row := make([]byte, n)
		for j:=0;j<n;j++{
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

func main() {
	fmt.Println(solveNQueens(8))
}