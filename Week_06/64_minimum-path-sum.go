package main

// dp 找出最小子结构，每个位置存放当前的最小步数
// dp[m][n] = min(dp[m-1][n], dp[m][n-1]) + nums[m][n]
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var dp [][]int
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		dp = append(dp, tmp)
	}

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] += dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] += dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {return a}
	return b
}



// dp 找出最小子结构，每个位置存放当前的最小步数
// dp[m][n] = min(dp[m-1][n], dp[m][n-1]) + nums[m][n]

// dp2 优化一维数组，不断右移
//func minPathSum(grid [][]int) int {
//	if len(grid) == 0 {
//		return 0
//	}
//	var dp = make([]int, len(grid[0]))
//	n := len(grid[0])
//	for i:=0; i<len(grid); i++ {
//		for j:=0; j<len(grid[0]); j++ {
//			if i == 0 && j == 0 {
//				dp[j] = grid[i][j]
//			} else if i == 0 {
//				dp[j] = dp[j-1] + grid[i][j]
//			} else if j == 0 {
//				dp[j] = dp[j] + grid[i][j]
//			} else {
//				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
//			}
//		}
//	}
//	return dp[n-1]
//}


// dp3 原数组，直接遍历操作 时间复杂度：O(N)，空间复杂度：O(1)
//func minPathSum(grid [][]int) int {
//	if len(grid) == 0 {
//		return 0
//	}
//	m, n := len(grid), len(grid[0])
//
//
//	for i:=0; i<len(grid); i++ {
//		for j:=0; j<len(grid[0]); j++ {
//			if i == 0 && j == 0 {
//				grid[i][j] = grid[i][j]
//			} else if i == 0 {
//				grid[i][j] = grid[i][j-1] + grid[i][j]
//			} else if j == 0 {
//				grid[i][j] = grid[i-1][j] + grid[i][j]
//			} else {
//				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
//			}
//		}
//	}
//	return grid[m-1][n-1]
//}