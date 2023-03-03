package code

func Execuate_51(n int) [][]string {
	return solveNQueens(n)
}

var ans_51 [][]string

func solveNQueens(n int) [][]string {
	//创建一个空棋盘
	chessboard := make([][]byte, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = '.'
		}
	}
	// fmt.Printf("board=%v\n", chessboard)
	ans_51 = make([][]string, 0)
	backtrace_51(chessboard, 0, n)
	return ans_51
}

func backtrace_51(chessboard [][]byte, row, n int) {
	if row == n {
		// fmt.Printf("board=%v, row=%d\n", chessboard, row)
		newboard := make([]string, 0)
		for _, v := range chessboard {
			newboard = append(newboard, string(v))
		}
		ans_51 = append(ans_51, newboard)
		return
	}

	for col := 0; col < n; col++ {
		//看是否能放Q
		if !isVaild_51(chessboard, row, col, n) {
			continue
		}
		chessboard[row][col] = 'Q'
		backtrace_51(chessboard, row+1, n)
		chessboard[row][col] = '.'
	}
}

func isVaild_51(chessboard [][]byte, row, col, n int) bool {
	//同列有没有
	for i := 0; i <= row; i++ {
		if chessboard[i][col] == 'Q' {
			return false
		}
	}
	//左斜上方有没有
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if chessboard[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	//右斜上方有没有
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if chessboard[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
