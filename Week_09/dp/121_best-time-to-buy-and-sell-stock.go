package main

// dp
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	// 0：用户手上不持股所能获得的最大利润，特指卖出股票以后的不持股，非指没有进行过任何交易的不持股
	// 1：用户手上持股所能获得的最大利润

	// 注意：因为题目限制只能交易一次，因此状态只可能从 1 到 0，不可能从 0 到 1
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i]) // 每次持有只能-当天的股价，没有历史积累
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp i-天数，k-允许交易次数（1），0、1表示是否持有 ;  只能完成一次交易
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, -prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
