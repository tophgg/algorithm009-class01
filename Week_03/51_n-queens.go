package main

// 解法1：回溯 使用map或slice存放行列、对角线的情况，按列循环对比
// 时间复杂度O(N!)，每个行列几乎都要对比n*(n-2)(n-4).... 空间复杂度O(N)，保存皇后的位置，对角线的位置
// 解法2：位运算
func solveNQueens(n int) [][]string {
	// terminator
	if n == 0 {
		return nil
	}
	// prepare end , data set
	res := make([][]int, 0)
	cols := make(map[int]bool, n)
	pies := make(map[int]bool, n)
	nas := make(map[int]bool, n)
	// dfs loop
	dfs([]int{}, n, cols, pies, nas, &res)
	// return result
	return generateResult(res, n)
}

func dfs(rows []int, n int, cols, pies, nas map[int]bool, res *[][]int) {
	// terminator
	row := len(rows)
	if row == n {
		tmp := make([]int, len(rows))
		copy(tmp, rows) // 由于rows每行都在引用，所以出rows时要新建一个临时变量
		*res = append(*res, tmp)
		return
	}
	// logic
	for col:=0; col<n; col++ {
		//当列，撇，捺中都不存在对应的值时，将其加入
		if !cols[col] && !pies[row+col] && !nas[row-col] {
			cols[col]=true
			pies[row+col]=true
			nas[row-col]=true
			// drill down
			dfs(append(rows,col), n, cols, pies, nas, res)
			// clear this row flag
			cols[col]=false
			pies[row+col]=false
			nas[row-col]=false
		}
	}
}
// 生成返回
func generateResult(res [][]int, n int) (result [][]string) {
	for _, v := range res {
		var s []string
		for _, val := range v{
			str := ""
			for i:=0; i<n; i++ {
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
	return result
}
