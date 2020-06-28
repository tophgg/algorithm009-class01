package main

// 解法1：暴力,双重循环模拟二维数组
// 解法2：dp，dp[i][j] = dp[i-1][j] + dp[i][j-1]  时间复杂度：O(m*n)，空间复杂度：O(m*n)
func uniquePaths(m int, n int) int {
	var dp = make([][]int, m, m)
	for i:=0; i<m; i++ {
		dp[i] = make([]int, n)
	}

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if i == 0 || j == 0 { // 将第一列、第一行置为1，其他格子根据左上相加得到相邻的步数
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 解法3：优化dp，一维数组只记录头上那一行，每一列用一个临时元素不断覆盖即可
//func uniquePaths(m int, n int) int {
//	var dp = make([]int, n)
//	// for i:=0; i<m; i++ {
//	//     dp[i] = make([]int, n)
//	// }
//
//	for i:=0; i<m; i++ {
//		for j:=0; j<n; j++ {
//			if j == 0 {
//				dp[j] = 1
//			} else {
//				dp[j] += dp[j-1]
//			}
//		}
//	}
//	return dp[n-1]
//}