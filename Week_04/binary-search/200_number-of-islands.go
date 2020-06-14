package main

// 解法1：dfs
// 解法2：bfs-dfs
// 解法3：并查集 flood fill
func numIslands(grid [][]byte) int {
	result := 0
	for i :=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if grid[i][j] == '1' {
				// 先将周围置0
				DFS(grid, i, j)
				result++
			}
		}
	}
	return result
}
var fx = [4]int{1, 0, -1, 0} // x轴 表示向上，第二步向左，第三步向上，第四步向右
var fy = [4]int{0, -1, 0, 1} // y轴
func DFS(grid [][]byte, x int, y int) {
	// 越界则退出， 或者当前节点为0
	if x < 0 || x >= len(grid) || y < 0 || y >=len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	// 向4个方向递归，使周围的点同化为0
	for i:=0; i<4; i++ {
		DFS(grid, x+fx[i], y+fy[i])
	}
}


// bfs-dfs 队列bfs标记

var fx = [4]int{1, 0, 0, -1}
var fy = [4]int{0, 1, -1, 0}
var row, col int

type Mark struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	// 标记位 记录
	m := make(map[Mark]bool)
	row = len(grid)
	if row == 0 {
		return 0
	}
	col = len(grid[0])
	count := 0
	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			_, marked := m[Mark{i,j}];
			if !marked && grid[i][j] == '1' {
				count++
				m[Mark{i,j}] = true
				queue := make([]int, 0)
				queue = append(queue, i, j)
				for len(queue) > 0 {
					x, y := queue[0], queue[1]
					queue = queue[2:]
					for k:=0; k<4; k++ {
						new_x := fx[k] + x
						new_y := fy[k] + y
						_, marked = m[Mark{new_x,new_y}]
						if new_x >= 0 && new_y >= 0 && new_x < row && new_y < col && !marked && grid[new_x][new_y] == '1' {
							queue = append(queue, new_x, new_y)
							m[Mark{new_x, new_y}] = true
						}
					}
				}
			}
		}
	}
	return count
}


// bfs-dfs
//var fx = [4]int{1, 0, 0, -1}
//var fy = [4]int{0, 1, -1, 0}
//var row, col int
//func numIslands(grid [][]byte) int {
//	row = len(grid)
//	if row == 0 {
//		return 0
//	}
//	col = len(grid[0])
//	count := 0
//	for i:=0; i<row; i++ {
//		for j:=0; j<col; j++ {
//			if grid[i][j] == '1' {
//				BFS(grid, i, j)
//				count++
//			}
//		}
//	}
//	return count
//}
//
//func BFS(grid [][]byte, x int, y int) {
//	queue := make([]int, 0)
//	queue = append(queue, x, y)
//	for len(queue) > 0 {
//		x, y = queue[0], queue[1]
//		queue = queue[2:]
//		for i:=0; i<4; i++ {
//			tmpX := x + fx[i]
//			tmpY := y + fy[i]
//			if tmpX >= 0 && tmpY >= 0 && tmpX < row && tmpY < col && grid[tmpX][tmpY] == '1' {
//				grid[tmpX][tmpY] = '0'
//				BFS(grid, tmpX, tmpY)
//			}
//		}
//	}
//}


var fx = [4]int{1, 0, 0, -1}
var fy = [4]int{0, 1, -1, 0}
var row, col int

type Mark struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	// 标记位 记录
	m := make(map[Mark]bool)
	row = len(grid)
	if row == 0 {
		return 0
	}
	col = len(grid[0])
	count := 0
	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			_, marked := m[Mark{i,j}];
			if !marked && grid[i][j] == '1' {
				count++
				m[Mark{i,j}] = true
				queue := make([]int, 0)
				queue = append(queue, i, j)
				for len(queue) > 0 {
					x, y := queue[0], queue[1]
					queue = queue[2:]
					for k:=0; k<4; k++ {
						new_x := fx[k] + x
						new_y := fy[k] + y
						_, marked = m[Mark{new_x,new_y}]
						if new_x >= 0 && new_y >= 0 && new_x < row && new_y < col && !marked && grid[new_x][new_y] == '1' {
							queue = append(queue, new_x, new_y)
							m[Mark{new_x, new_y}] = true
						}
					}
				}
			}
		}
	}
	return count
}