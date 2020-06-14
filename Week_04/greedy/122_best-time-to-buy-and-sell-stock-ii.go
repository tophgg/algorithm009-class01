package main

// 解法1：贪心 每天买入，
// 解法2：dp
func maxProfit(prices []int) int {
	sum := 0
	perMoney := prices[0]
	for i:=1; i<len(prices); i++ {
		if prices[i] > perMoney {
			sum += prices[i] - perMoney
		}
		perMoney = prices[i]

	}
	return sum
}