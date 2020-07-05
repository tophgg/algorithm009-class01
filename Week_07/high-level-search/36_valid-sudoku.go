package main

// 一次迭代，剪枝，使用map存储，分别存储行、列、和3*3块
// 块的公式是i/3 * 3 + j/3
func isValidSudoku(board [][]byte) bool {

	for i:=0; i<9; i++ {
		row := map[byte]bool{}
		col := map[byte]bool{}
		block := map[byte]bool{}
		for j:=0; j<9; j++ {
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}

			if board[j][i] != '.' {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}

			// i/3 * 3 + j/3
			r := i%3 * 3 + j%3
			c := i/3 * 3 + j/3
			if board[r][c] != '.' {
				if block[board[r][c]] {
					return false
				}
				block[board[r][c]] = true
			}
		}
	}
	return true
}