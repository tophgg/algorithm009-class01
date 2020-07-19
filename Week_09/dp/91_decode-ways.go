package main

// dp，每次移动，相当于可以统计1个 或统计2个
// 注意边界，当前为0的情况， 开dp数组预留+1
// 当前两个数为10~26时，dp[i+1] = dp[i-1] + dp[i] 走1步或走2步
// 当为20和10时，相当于dp[i+1] = dp[i-1] 因为这2个数只能看作一个数编码
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else {
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
				dp[i+1] = dp[i-1] + dp[i]
			} else {
				dp[i+1] = dp[i]
			}
		}
	}
	return dp[len(s)]
}
