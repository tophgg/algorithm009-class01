package main

// 1.dp，注意可跨越，所以需要对前j个元素进行比较，dp[i] = max(dp[j]+1, dp[i])；
// 时间复杂度O(N^2)，空间复杂度O(N)
// 2.二分 时间复杂度O(logN)，空间复杂度O(N)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				maxLen = max(maxLen, dp[i])
			}
		}
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0
	top := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		left := bindSearch(nums[i], top, length)
		top[left] = nums[i]
		if left == length {
			length++
		}
	}
	return length
}

func bindSearch(target int, top []int, length int) int {
	l, r := 0, length-1
	for l <= r {
		mid := l + (r-l)>>1
		if target > top[mid] {
			l++
		} else if target < top[mid] {
			r--
		} else {
			return mid
		}
	}
	return l
}
