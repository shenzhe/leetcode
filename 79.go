package code

func Execuate_79(board [][]byte, word string) bool {
	return exist(board, word)

}

func exist(board [][]byte, word string) bool {
	wLen, rowLen, colLen := len(word), len(board[0])-1, len(board)-1
	for i := 0; i <= colLen; i++ {
		for j := 0; j <= rowLen; j++ {
			if board[i][j] != word[0] {
				continue
			}
			if traverse_79(board, j, i, 0, wLen, rowLen, colLen, word) {
				return true
			}
		}
	}
	return false
}

func traverse_79(board [][]byte, x, y, idx, wLen, rowLen, colLen int, word string) bool {

	if idx == wLen {
		return true
	}

	if x < 0 || x > rowLen || y < 0 || y > colLen {
		return false
	}

	// fmt.Printf("x=%d, y=%d, s=%c, w=%c rowLen=%d, idx=%d, colLen=%d\n", x, y, board[y][x], word[idx], idx, rowLen, colLen)

	if board[y][x] != word[idx] {
		return false
	}

	//标记已访问过，且可恢复，省掉一个额外的visit数组
	board[y][x] = 0 - board[y][x]
	//上
	if traverse_79(board, x, y-1, idx+1, wLen, rowLen, colLen, word) {
		return true
	}
	//下
	if traverse_79(board, x, y+1, idx+1, wLen, rowLen, colLen, word) {
		return true
	}
	//左
	if traverse_79(board, x-1, y, idx+1, wLen, rowLen, colLen, word) {
		return true
	}
	//右
	if traverse_79(board, x+1, y, idx+1, wLen, rowLen, colLen, word) {
		return true
	}
	//本轮遍历完面，还原数字，方便下轮遍历
	board[y][x] = 0 - board[y][x]
	return false
}
