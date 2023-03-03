package code

func Execuate_52(n int) int {
	return totalNQueens(n)
}

var count_52 int

func totalNQueens(n int) int {
	//构建棋盘
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	count_52 = 0
	//穷举棋盘，从第一行开始摆皇后
	traverse_52(board, n, 0)
	return count_52
}

func traverse_52(board [][]int, n, row int) {
	//如果遍历到了最后一行，说明是正确角
	if row == n {
		count_52++
		return
	}

	//分别在相应的格子摆皇后
	for col := 0; col < n; col++ {
		//判断此行此位置是否能摆皇后,如果不能摆，跳过
		if !isValid_52(board, row, col, n) {
			continue
		}
		//摆上皇后
		board[row][col] = 1
		//进入下一行
		traverse_52(board, n, row+1)
		//row行col列摆放检测完毕，恢复空盘，进行下一个位置
		board[row][col] = 0
	}
}

func isValid_52(board [][]int, row, col, n int) bool {
	//同一列的位置的判断
	for i := 0; i <= row; i++ {
		//同一列已摆了皇后，不能摆放皇后
		if board[i][col] == 1 {
			return false
		}
	}

	//左上方斜列不能有(row--, col--的位置)
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}
	//右上方斜列不能有(row--, col++的位置)
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if board[i][j] == 1 {
			return false
		}
		i--
		j++
	}

	return true
}
