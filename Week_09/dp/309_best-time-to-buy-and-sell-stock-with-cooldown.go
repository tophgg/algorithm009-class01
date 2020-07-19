package main

// dp思路 j表示当前1买入，0卖出，冷冻 dp[i][1] = max(dp[i-1][0])
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = Max(dp[0][1]+prices[1], dp[0][0])
	dp[1][1] = Max(dp[0][1], -prices[1])
	for i := 2; i < length; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-2][0]-prices[i]) // 冷冻期指前一天的未持有才能购买
	}
	return dp[length-1][0]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dp思路 j表示当前1买入，0卖出，冷冻 dp[i][1] = max(dp[i-1][0])
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	// 初始化dp20为2天前未购买，dp10代表昨天没买，dp11代表昨天买了
	dp20 := 0
	dp10 := 0
	dp11 := -prices[0]
	for i := 1; i < length; i++ {
		// 昨天没买记录下来，当做下一次的前天没买的值
		tmp := dp10
		// 更新今天的没买和买的情况
		dp10 = Max(dp10, dp11+prices[i])
		dp11 = Max(dp11, dp20-prices[i])
		dp20 = tmp
	}
	return dp10
}
