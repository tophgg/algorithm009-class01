package main

// 解法1：dfs
// 解法2：bfs-dfs-dfs
// 解法3：并查集 flood fill
func numIslands(grid [][]byte) int {
	result := 0
	for i :=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if grid[i][j] == '1' {
				DFS(grid, i, j)
				result++
			}
		}
	}
	return result
}
var fx = [4]int{-1, 1, 0, 0} // x轴 表示向上，第二步向左，第三步向上，第四步向右
var fy = [4]int{0, 0, -1, 1} // y轴
// var fx = [4]int{1, 0, -1, 0} // x轴 表示向上，第二步向左，第三步向上，第四步向右
// var fy = [4]int{0, -1, 0, 1} // y轴
func DFS(grid [][]byte, x int, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >=len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	for i:=0; i<4; i++ {
		DFS(grid, x+fx[i], y+fy[i])
	}
}