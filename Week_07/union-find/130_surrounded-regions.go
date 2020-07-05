package main

// dfs
// 并查集 一般用一维数组来记录，方便查找parent，所以需要将二维坐标转为一维坐标 x * row + y
var Parent = make([]int, 0)
var Count = 0
// var Size = []int{}
var col, row = 0, 0
func Init(n int) {
	Count = n
	Parent = make([]int, n)
	// Size = make([]int, n)
	for i:=0; i<n; i++ {
		Parent[i] = i
		// Size[i] = i
	}
}

// 返回某个节点p的根节点，主要复杂度，即待优化api
// 如果将这个树高度减为2，可以减少遍历次数
func Find(p int) int {
	// 根节点的parent[p] = p
	for p != Parent[p] {
		// 路径压缩，降为O(1)
		Parent[p] = Parent[Parent[p]]
		p = Parent[p]
	}
	return p
}

// 合并2个集合
// 可做平衡性优化，小的树接到大的树上，用size记录每个树的高度
func Union(p, q int) {
	rootP := Find(p)
	rootQ := Find(q)
	if rootP == rootQ {
		return
	}
	// 以其中一个集合为顶点
	Parent[rootP] = rootQ
	// if Size[rootP] > Size[rootQ] {
	// 	Parent[rootQ] = rootP
	// 	Size[rootP] += Size[rootQ]
	// } else {
	// 	Parent[rootP] = rootQ
	// 	Size[rootQ] += Size[rootP]
	// }
	Count--
}

func IsConnected(p, q int) bool {
	p = Find(p)
	q = Find(q)
	return p == q
}
func node(x, y int) int {
	return x * col + y
}
func solve(board [][]byte)  {
	row = len(board)
	if row == 0 { return }
	col = len(board[0])
	Init(row * col + 1)
	dummyNode := row * col

	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			if board[i][j] == 'O' {
				// 边界上的O和dummyNode连通
				if i == 0 || i==row-1 || j==0 || j==col-1 {
					Union(node(i,j), dummyNode)
				} else {
					// 把上下左右合并成一个连通区域
					if i>0 && board[i-1][j] == 'O' {
						Union(node(i,j), node(i-1,j))
					}
					if i<row-1 && board[i+1][j] == 'O' {
						Union(node(i,j), node(i+1,j))
					}
					if j>0 && board[i][j-1] == 'O' {
						Union(node(i,j), node(i,j-1))
					}
					if j<col-1 && board[i][j+1] == 'O' {
						Union(node(i,j), node(i,j+1))
					}
				}
			}
		}
	}
	// 判断是否连通， 只需该内核
	for i:=1; i<row; i++ {
		for j:=1; j<col; j++ {
			if !IsConnected(node(i,j), dummyNode) {
				board[i][j] = 'X'
			}
		}
	}
}



//dfs
//func dfs(board [][]byte, i, j int) {
//	// 边界判断
//	// board[i][j] == '#': 说明已经搜索过
//	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
//		return
//	}
//	// 和边界联通的O，先替换成#， 后续换回O
//	board[i][j] = '#'
//	dfs(board, i+1, j)
//	dfs(board, i-1, j)
//	dfs(board, i, j+1)
//	dfs(board, i, j-1)
//}
//
//func solve(board [][]byte) {
//	if board == nil || len(board) == 0 {
//		return
//	}
//	// 先搜索并标记
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[0]); j++ {
//			// 从边缘开始搜索, 判断是否在边上，只从边上开始，如果边上都是X，那么里面的所有O都换成X就行
//			// 如果边上有，那么之后就会被替换成#暂存，之后会被恢复成O，意味着不换成X
//			isEdge := i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1
//			if isEdge && board[i][j] == 'O' {
//				dfs(board, i, j)
//			}
//		}
//	}
//	// 再替换，这里的O都应该是被包围的
//	// # 是和边界联通的，即不被包围的
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[0]); j++ {
//			if board[i][j] == 'O' {
//				board[i][j] = 'X'
//			}
//			if board[i][j] == '#' {
//				board[i][j] = 'O'
//			}
//		}
//	}
//}
