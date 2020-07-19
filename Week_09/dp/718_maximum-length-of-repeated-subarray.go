package main

// 解法1：dp   dp[i][j] 为当前A,B长度下的公共数组； 时间复杂度O(N*M),空间复杂度O(N*M)， ##利用滚动数组可以优化到 O(min(N,M))。
// dp[0][0] = 当前i,j > 0 判断 A[i] == A[j] => dp[i][j] = dp[i-1][j-1] + 1，记录历史最大值   判断A[i] != A[j] => dp[i][j] = 0  最后判断是否历史最大值
// 解法2：滑动窗口
// 时间复杂度O(N+M)*MIN(N,M),空间复杂度O(1)
// 解法3：二分查找+哈希
// 时间复杂度O((N+M)log(MIN(N,M)))，空间复杂度O(N)
func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		dp[i] = make([]int, len(B)+1)
	}
	res := 0
	for i := 1; i < len(A)+1; i++ {
		for j := 1; j < len(B)+1; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			res = max(dp[i][j], res)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 滑动窗口
func findLength(A []int, B []int) int {
	lenA := len(A)
	lenB := len(B)
	res := 0
	maxLen := 0
	length := 0
	// B不动A动
	for i := 0; i <= lenB; i++ {
		length = min(lenA, lenB-i)             // 当前窗口的长度，B固定不动，每次B向右移动
		maxLen = getMaxLen(A, B, 0, i, length) // 窗口从A[0] ~ B[i]
		res = max(res, maxLen)
	}
	// A不动B动
	for i := 0; i <= lenA; i++ {
		length = min(lenB, lenA-i)             // 当前窗口的长度，A固定不动，每次A向右移动
		maxLen = getMaxLen(A, B, i, 0, length) // 窗口从B[0] ~ A[i]
		res = max(res, maxLen)
	}
	return res
}

// 遍历窗口内A,B的元素，只要有相等的就++，有一个不相等就清零
func getMaxLen(A []int, B []int, astart int, bstart int, windowLen int) int {
	res := 0
	subLen := 0
	for i := 0; i < windowLen; i++ {
		if A[astart+i] == B[bstart+i] {
			subLen++
		} else {
			subLen = 0
		}
		res = max(res, subLen)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
