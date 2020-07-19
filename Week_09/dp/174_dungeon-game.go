package main

import "math"

// dp 二维，因为正向推导时，受路程中最小值和总路径和 两个维度的因素影响，无法正常推导，故从后往前推导
// 时间复杂度O(M*N),空间复杂度O(M*N)

func calculateMinimumHP(dungeon [][]int) int {
	row := len(dungeon)
	col := len(dungeon[0])
	dp := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		dp[i] = make([]int, col+1)
		for j := 0; j < col+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// 多余边界给作为i+1，j+1的最低判断值，即至少有1滴血
	dp[row][col-1], dp[row-1][col] = 1, 1
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			minValue := min(dp[i+1][j], dp[i][j+1])   // 向右/向下的当前最小路径
			dp[i][j] = max(minValue-dungeon[i][j], 1) // dp存最大的路径需要血量值  1滴血走完中点or实际需要抵扣的值
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
