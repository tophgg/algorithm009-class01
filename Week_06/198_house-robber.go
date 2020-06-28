package main


// dp1：二维数组，记录每户的金额+是否有抢
// dp[i][0] = max(dp[i-1][1], dp[i-1][0])  // 当前房子不偷 = 前一个房子偷和不偷的最大值
// dp[i][1] = dp[i-1][0]+nums[i]           // 当前房子偷 = 前一个房子不偷+当前房子的值
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var dp = make([][]int, len(nums))
	dp[0] = append(dp[0], 0, nums[0])
	for i:=1; i<len(nums); i++ {
		// dp[i][0] 存放不偷房子的最大值，dp[i][1]存放偷当前房子的值
		dp[i] = append(dp[i], max(dp[i-1][1], dp[i-1][0]), dp[i-1][0] + nums[i])
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}


// dp2: 一维数组，每个值存放当前抢的最大值
// dp3: 用一个maxValue存放当前最大值，不断交替即可
//func rob(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	var dp = make([]int, len(nums))
//	dp[0] = nums[0]
//	dp[1] = max(nums[0], nums[1])
//	for i:=2; i<len(nums); i++ {
//		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
//	}
//	return dp[len(nums)-1]
//}


// dp3: 用一个maxValue存放当前最大值，不断交替即可
//func rob(nums []int) int {
//	preVal := 0
//	curVal := 0
//	lastVal := 0
//	for i:=0; i<len(nums); i++ {
//		lastVal = max(curVal, preVal+nums[i])
//		preVal, curVal = curVal, lastVal
//	}
//	return curVal
//}