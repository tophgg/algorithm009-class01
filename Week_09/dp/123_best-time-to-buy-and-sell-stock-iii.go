package main

import "math"

// dp三维 i表示第几天，j表示是否持有,k表示已完成交易几天，
// dp[0][0][0] = 0
// dp[0][1][0] = -prices[0]   // 注意初始化为-prices[0]或负无穷

// 从未持有
// dp[i][0][0] = 0
// 未持有，但交易过1次，可能是今天交易或昨天交易
// dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][1][0]+prices[i])
// 未持有，交易过2次，可能是昨天已经交易2次，或者昨天交易了1次并将持有的卖出
// dp[i][0][2] = max(dp[i-1][0][2], dp[i-1][1][1]+prices[i])
// 持有，未交易，前一天持有且未交易，或者是昨天未持有，今天交易
// dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0] - prices[i])
// 持有，交易过1次；可能是昨天就持有，或者昨天未持有，今天买入
// dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][1] - prices[i])
// 持有，交易过2次，不存在
// dp[i][1][2] = -1
func maxProfit(prices []int) int {
	dp := make([][][]int, len(prices))
	length := len(prices)
	if length == 0 {
		return 0
	}
	for i := 0; i < length; i++ {
		dp[i] = make([][]int, 2)
		for k := 0; k < 2; k++ {
			dp[i][k] = make([]int, 3)
		}
	}
	dp[0][0][0] = 0
	dp[0][1][0] = -prices[0]
	dp[0][1][1] = -prices[0]
	dp[0][1][2] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][1][0]+prices[i])
		dp[i][0][2] = max(dp[i-1][0][2], dp[i-1][1][1]+prices[i])
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0]-prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][1]-prices[i])
		dp[i][1][2] = math.MinInt32
	}
	//
	return max(dp[length-1][0][2], max(dp[length-1][0][1], dp[length-1][0][0]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp三维搞起
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	length := len(prices)
	// dp0为未持有未交易，dp1表示持有未交易，
	// dp2表示未持有已交易1次，dp3表示持有已交易1次，
	// dp4表示未持有已交易2次
	dp0, dp1, dp2, dp3, dp4 := 0, -prices[0], -prices[0], -prices[0], -prices[0]
	for i := 1; i < length; i++ {
		dp1 = max(dp1, dp0-prices[i])
		dp2 = max(dp2, dp1+prices[i])
		dp3 = max(dp3, dp2-prices[i])
		dp4 = max(dp4, dp3+prices[i])
	}
	return max(dp2, dp4)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
