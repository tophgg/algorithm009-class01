package main

// dp 奇数是010 -> 011 前一个+1；偶数是010 -> 100，2的倍数相同；base case0只有0
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i:=1; i<=num; i++ {
		if i & 1 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}