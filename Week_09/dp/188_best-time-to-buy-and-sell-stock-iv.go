package main

import "math"

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}
	length := len(prices)
	if k > length/2 {
		return maxProfit2(prices)
	}
	dp := make([][][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}
	dp[0][0][0] = 0
	for p := 0; p <= k; p++ {
		dp[0][1][p] = -prices[0]
	}
	maxPro := 0
	for i := 1; i < length; i++ {
		dp[i][0][0] = dp[i-1][0][0]
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0]-prices[i])
		for p := 0; p <= k; p++ {
			if p == 0 {
				dp[i][0][p] = dp[i-1][0][p]
				dp[i][1][p] = max(dp[i-1][1][p], dp[i-1][0][p]-prices[i])
			} else {
				dp[i][1][p] = max(dp[i-1][1][p], dp[i-1][0][p]-prices[i])
				dp[i][0][p] = max(dp[i-1][0][p], dp[i-1][1][p-1]+prices[i])
				maxPro = max(maxPro, dp[i][0][p])
			}
		}
	}
	return maxPro
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp一维优化
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}
	length := len(prices)
	if k > length/2 {
		return maxProfit2(prices)
	}
	dp0, dp1 := make([]int, k+1), make([]int, k+1)
	for p := 0; p <= k; p++ {
		dp0[p] = 0
		dp1[p] = math.MinInt64
	}
	for _, v := range prices {
		for p := 1; p <= k; p++ {
			dp0[p] = max(dp0[p], dp1[p]+v)
			dp1[p] = max(dp1[p], dp0[p-1]-v)
		}
	}
	return dp0[k]
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, dp0-prices[i])
	}
	return dp0
}
