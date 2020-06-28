package main

// 解法1：暴力，时间复杂度：O(m*n*min(m,n)^2)，空间复杂度：O(1)
// 解法2：dp，遍历二维数组，维护dp数组，为左上3个点的最小值+1，为当前点的最长正方形的边，dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i:=0; i<len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j:=0; j<len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i:=1; i<len(matrix); i++ {
		for j:=1; j<len(matrix[0]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
func min(a, b int) int {
	if a < b {return a}
	return b
}