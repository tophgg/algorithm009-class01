package main

// 解法1：dfs 着色，8个方向
// 解法2：bfs-dfs-dfs
var d = [8][2]int{{1,0},{-1,0},{0,1},{0,-1},{1,1},{1,-1},{-1,1},{-1,-1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	a, b := click[0], click[1]
	// 如果中雷，标记结束
	if board[a][b] == 'M' {
		board[a][b] = 'X'
	} else if board[a][b] == 'E' {
		// 踩空后，要判断dfs遍历周围8个节点的状态
		m, n := len(board), len(board[0])
		var f func(int, int); f = func(i, j int) {
			c := byte('0')
			// 统计周围的雷数
			for _, di := range d {
				x, y := i + di[0], j + di[1]
				if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'M' {
					c++
				}
			}
			// 如果有雷，更新为雷的个数
			if c > '0' {
				board[i][j] = c
			} else {
				// 如果为周围没雷，更新为空，然后继续统计周围为E的个数
				board[i][j] = 'B'
				for _, di := range d {
					x, y := i + di[0], j + di[1]
					if x >= 0 && x < m && y>=0 && y < n && board[x][y] == 'E' {
						f(x, y)
					}
				}
			}
		}
		f(a, b)
	}

	return board
}