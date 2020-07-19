package main

// dp二维
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	length := len(prices)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-fee+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[length-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp 一次
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	length := len(prices)
	// dp0表示不持有，dp1表示持有
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < length; i++ {
		dp0 = max(dp0, dp1+prices[i]-fee)
		dp1 = max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
