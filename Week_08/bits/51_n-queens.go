package main
//时间复杂度：O(N!)
//空间复杂度：O(N) 位运算的话
func solveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}
	result := make([][]int, 0)
	solve([]int{}, &result, n, 0, 0, 0)
	return dealWithResult(result)
}

// 分别用4个二进制位来记录每行每列的访问情况
func solve(rows []int, result *[][]int, n int, cols, pies, nas int) {
	row := len(rows)
	if row >= n {
		//*result = append(*result, )
		tmp := make([]int, len(rows))
		copy(tmp, rows)
		*result = append(*result, tmp)
		return
	}
	// 32位n，比如8皇后只需要用到低8位，需要对高位取0处理，左移8位再取反 获得全1
	// 取出当前行可放置皇后的格子 cols | pies | nas
	// bits记录的是每一次遍历当前行的棋盘，1表示可放置；0标识不可放置
	bits := ^(cols | pies | nas) & (1<<uint(n)-1)
	for bits != 0 {
		// 每次取最右边的低位1 用于测试放入皇后置为0   （每次循环取最低位1 并将其置为0）
		last := bits & (-bits)
		col := 0
		// 直到test的位置为可放入，即1
		for test:=1; last & test == 0; test <<= 1 {
			col++
		}
		// | 或者 ^
		solve(append(rows, col), result, n, cols | last, (pies | last) << 1, (nas | last) >> 1)
		// 最右边的低位1置为0表示已选
		// bits ^= last
		bits = bits & (bits-1)
	}
}


func dealWithResult(res [][]int) (result [][]string) {
	for _, v := range res {
		var s []string
		for _, val := range v {
			str := ""
			for i := 0; i < len(v); i++ {
				if i == val {
					str += "Q"
				} else {
					str += "."
				}
			}
			s = append(s, str)
		}
		result = append(result, s)
	}
	return
}

