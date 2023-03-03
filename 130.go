package code

import "fmt"

func Execuate_130(board [][]byte) {
	solve(board)
}

func Execuate_130_2(board [][]byte) {
	solve_2(board)
}

func Execuate_130_3(board [][]byte) {
	solve_3(board)
}

func solve(board [][]byte) {
	xLen, yLen := len(board[0])-1, len(board)-1
	fmt.Printf("xlen=%d, yLen=%d\n", xLen, yLen)
	visited := make(map[string]bool)
	onpath := make(map[string][][2]int)
	flag := make(map[string]bool)
	for y := 0; y <= yLen; y++ {
		for x := 0; x <= xLen; x++ {
			// fmt.Printf("loop：x=%d, y=%d\n", x, y)
			if board[y][x] == 'X' {
				continue
			}
			vk := fmt.Sprintf("%d_%d", y, x)
			if _, ok := visited[vk]; ok {
				continue
			}
			onpath[vk] = make([][2]int, 0)
			// fmt.Printf("spath=%v, x=%d, y=%d, flag=%t\n", onpath, x, y, flag[vk])
			dfs_130(board, x, y, xLen, yLen, visited, onpath, vk, flag)
			// fmt.Printf("epath=%v, x=%d, y=%d, flag=%t\n", onpath, x, y, flag[vk])
		}
	}

}

func dfs_130(board [][]byte, x, y, xLen, yLen int, visited map[string]bool, path map[string][][2]int, pvk string, flag map[string]bool) {
	//超出边界
	if x < 0 || y < 0 || x > xLen || y > yLen {
		return
	}

	//是X跳过
	if board[y][x] == 'X' {
		return
	}

	vk := fmt.Sprintf("%d_%d", y, x)
	if _, ok := visited[vk]; ok {
		return
	}
	visited[vk] = true
	path[pvk] = append(path[pvk], [2]int{y, x})
	fmt.Printf("path1=%v, x=%d, y=%d, flag=%t\n", path[pvk], x, y, flag[pvk])
	//标记为X
	if !flag[pvk] {
		board[y][x] = 'X'
	}
	if !flag[pvk] && (x == 0 || y == 0 || x == xLen || y == yLen) {
		flag[pvk] = true
		board[y][x] = 'O'
		if len(path) > 0 { //还原为O
			for _, v := range path[pvk] {
				board[v[0]][v[1]] = 'O'
			}
		}
	}

	dfs_130(board, x-1, y, xLen, yLen, visited, path, pvk, flag) //左
	dfs_130(board, x+1, y, xLen, yLen, visited, path, pvk, flag) //右
	dfs_130(board, x, y-1, xLen, yLen, visited, path, pvk, flag) //上
	dfs_130(board, x, y+1, xLen, yLen, visited, path, pvk, flag) //下
	// fmt.Printf("path2=%v, x=%d, y=%d, flag=%t\n", path[pvk], x, y, flag[pvk])
}

// 遍历四边
func solve_2(board [][]byte) {
	xLen, yLen := len(board[0])-1, len(board)-1
	fmt.Printf("xlen=%d, yLen=%d\n", xLen, yLen)
	visited := make(map[string]bool)
	for y := 0; y <= yLen; y++ {
		for x := 0; x <= xLen; x++ {
			// fmt.Printf("loop：x=%d, y=%d\n", x, y)
			if board[y][x] == 'X' {
				continue
			}
			if !(x == 0 || y == 0 || x == xLen || y == yLen) {
				continue
			}
			vk := fmt.Sprintf("%d_%d", y, x)
			if _, ok := visited[vk]; ok {
				continue
			}
			// fmt.Printf("spath=%v, x=%d, y=%d, flag=%t\n", onpath, x, y, flag[vk])
			dfs_130_2(board, x, y, xLen, yLen, visited)
			// fmt.Printf("epath=%v, x=%d, y=%d, flag=%t\n", onpath, x, y, flag[vk])
		}
	}

	for y := 0; y <= yLen; y++ {
		for x := 0; x <= xLen; x++ {
			if board[y][x] == 'X' {
				continue
			}
			if board[y][x] == '0' {
				board[y][x] = 'X'
				continue
			}
			if board[y][x] == '#' {
				board[y][x] = 'O'
				continue
			}
		}
	}

}

func dfs_130_2(board [][]byte, x, y, xLen, yLen int, visited map[string]bool) {
	//超出边界
	if x < 0 || y < 0 || x > xLen || y > yLen {
		return
	}

	//是X跳过
	if board[y][x] == 'X' {
		return
	}

	vk := fmt.Sprintf("%d_%d", y, x)
	if _, ok := visited[vk]; ok {
		return
	}
	visited[vk] = true

	//标记为#
	board[y][x] = '#'
	dfs_130_2(board, x-1, y, xLen, yLen, visited) //左
	dfs_130_2(board, x+1, y, xLen, yLen, visited) //右
	dfs_130_2(board, x, y-1, xLen, yLen, visited) //上
	dfs_130_2(board, x, y+1, xLen, yLen, visited) //下
	// fmt.Printf("path2=%v, x=%d, y=%d, flag=%t\n", path[pvk], x, y, flag[pvk])
}

func solve_3(board [][]byte) {
	xLen, yLen := len(board[0]), len(board)
	// fmt.Println(xLen, yLen)
	uf := NewUnionFind(xLen*yLen + 1)

	dummy := xLen * yLen
	// fmt.Printf("count=%d, dummy=%d, parent=%v\n", uf.Count(), dummy, uf.GetParent())
	//2d => 1d x * xLen + y
	for i := 0; i < xLen; i++ {
		if board[0][i] == 'O' {
			uf.Union(twoToOne(i, 0, xLen), dummy)
		}
		if board[yLen-1][i] == 'O' {
			uf.Union(twoToOne(i, yLen-1, xLen), dummy)
		}
	}
	// fmt.Printf("count2=%d, parent=%v\n", uf.Count(), uf.GetParent())
	for j := 0; j < yLen; j++ {
		if board[j][0] == 'O' {
			uf.Union(twoToOne(0, j, xLen), dummy)
		}

		if board[j][xLen-1] == 'O' {
			uf.Union(twoToOne(xLen-1, j, xLen), dummy)
		}
	}
	// fmt.Printf("count3=%d, parent=%v\n", uf.Count(), uf.GetParent())

	for y := 1; y < yLen-1; y++ { //y
		for x := 1; x < xLen-1; x++ { //x
			if board[y][x] == 'X' {
				continue
			}
			//左 = x-1
			if board[y][x-1] == 'O' {
				uf.Union(twoToOne(x-1, y, xLen), twoToOne(x, y, xLen))
			}
			//右 = x+1
			if board[y][x+1] == 'O' {
				uf.Union(twoToOne(x+1, y, xLen), twoToOne(x, y, xLen))
			}
			//上 = y-1
			if board[y-1][x] == 'O' {
				uf.Union(twoToOne(x, y-1, xLen), twoToOne(x, y, xLen))
			}
			//下 = y+1
			if board[y+1][x] == 'O' {
				uf.Union(twoToOne(x, y+1, xLen), twoToOne(x, y, xLen))
			}

		}
	}

	// fmt.Printf("count4=%d, parent=%v\n", uf.Count(), uf.GetParent())
	//所有不与dummy连通的O都替换成X
	for y := 1; y < yLen-1; y++ { //y
		for x := 1; x < xLen-1; x++ { //x
			if !uf.Connected(dummy, twoToOne(x, y, xLen)) {
				board[y][x] = 'X'
			}
		}
	}
}

func twoToOne(x, y, xLen int) int {
	return y*xLen + x
}
