package main

// dp dp[i] = min(dp[i-1], dp[i-2]) + cost[i] 到目前为止的步数
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[len(cost)-1] = cost[len(cost)-1]
	dp[len(cost)-2] = cost[len(cost)-2]
	for i := len(cost) - 3; i >= 0; i-- {
		dp[i] = min(dp[i+1], dp[i+2]) + cost[i]
	}
	return min(dp[0], dp[1])
}

// dp一维 dp[i] = min(dp[i-1], dp[i-2]) + cost[i] 到目前为止的步数
func minCostClimbingStairs(cost []int) int {
	dp1, dp2 := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		dp0 := min(dp1, dp2) + cost[i]
		dp2 = dp1
		dp1 = dp0
	}
	return min(dp1, dp2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
