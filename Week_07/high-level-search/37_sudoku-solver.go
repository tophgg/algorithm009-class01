package main

// dfs + 回溯 可优化的点是，先申请数组，将数组填满已经有数字的点，再从最少选择的行开始做
// A*搜索，启发式搜索
// 解数独
func solveSudoku(board [][]byte)  {
	backstrace(board, 0, 0)
}

func backstrace(board [][]byte, row int, col int) bool {
	if row == 9 { return true }
	if col == 9 { // 换行
		return backstrace(board, row+1, 0)
	}
	if board[row][col] != '.' { // 跳过已有数字的格子
		return backstrace(board, row, col+1)
	}
	// 当前层逻辑
	for i:='1'; i<='9'; i++ {
		if isValid(board, row, col, byte(i)) {
			board[row][col] = byte(i) // 写入当前数字
			if backstrace(board, row, col+1) { // 下探下一层，如果都成功就返回true
				return true
			}
			board[row][col] = '.' // 复原，下一层失败，恢复原来的.
		}
	}
	return false
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i:=0; i<9; i++  {
		if board[row][i] == c {
			return false
		}
		if board[i][col] == c {
			return false
		}
		if board[row/3*3+i/3][col/3*3+i%3] == c {
			return false
		}
	}
	return true
}