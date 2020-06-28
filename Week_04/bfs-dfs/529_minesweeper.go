package main

// 扫雷游戏
var d = [8][2]int{{1,0},{0,1},{-1,0},{0,-1},{1,1},{1,-1},{-1,1},{-1,-1}}

// 解法1：bfs-dfs-dfs+dfs
func updateBoard(board [][]byte, click []int) [][]byte {
	a, b := click[0], click[1]
	// 先判断是否终点
	if board[a][b] == 'M' {
		board[a][b] = 'X'
	} else if board[a][b] == 'E' {
		m, n := len(board), len(board[0])
		var f func(int, int); f = func(i, j int) {
			c := byte('0')
			// 判断当前节点周围8个节点是否有雷，记录数字，bfs遍历每一层
			for _, di := range d {
				x, y := di[0] + i, j + di[1]
				if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'M' {
					c++
				}
			}
			// 如果有，则更新当前节点为雷数
			if c > '0' {
				board[i][j] = c
			} else {
				// 如果未空，则更新当前节点为空块，并且查找附近8个节点为E的节点，如果有节点为E则dfs遍历该节点
				board[i][j] = 'B'
				for _, di := range d {
					x, y := di[0] + i, di[1] + j
					if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'E' {
						// 下探
						f(x,y)
					}
				}
			}
		}
		f(a,b)
	}
	return board
}