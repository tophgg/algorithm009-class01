package main

// 待优化
var Parent = []int{}
var row, col = 0, 0
var Count = 0
func numIslands(grid [][]byte) int {
	col = len(grid)
	if col == 0 {return 0}
	row = len(grid[0])
	if row == 0 {return 0}
	dummyNode := col * row
	Count := col * row + 1
	Parent = make([]int, Count)

	for i:=0; i<Count; i++ {
		Parent[i] = i
	}
	// 定义4个方向的坐标值
	directions := [2][2]int{{0,1}, {1,0}} // 只判断右边和下面
	for i:=0; i<col; i++ {
		for j:=0; j<row; j++ {
			if grid[i][j] == 0 {
				union(getIndex(i,j), dummyNode)
			}
			if grid[i][j] == 1 {
				for _, d := range directions {
					nc, nr := i+d[0], j+d[j]
					if nc>=0 && nr >=0 && nc<col && nr<row && grid[nr][nc] == 1 {
						union(getIndex(i,j), getIndex(nc, nr))
					}
				}
			}
		}
	}
	return Count-1
}

func getIndex(x, y int) int {
	return x * col + y
}

func find(p int) int {
	for p != Parent[p] {
		Parent[p] = Parent[Parent[p]]
		p = Parent[p]
	}
	return p
}

func union(p, q int) {
	rootP := find(p)
	rootQ := find(q)
	if rootP == rootQ {
		return
	}
	Parent[rootP] = rootQ
	Count -=1
}