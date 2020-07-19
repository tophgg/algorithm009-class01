package main

// dp LCS 穷举+剪枝 遍历2边的子串，同时比较最后一个字符看是否相同
func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			// 当前遍历的字符相等，则取上一个公共值（未到当前lcs字符串之前的值）加1
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 不相等，则等于相邻当前对比获得的最大的公共数
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// dp LCS 穷举+剪枝 遍历2边的子串，同时比较最后一个字符看是否相同
func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	dp := make([]int, l2+1)
	tmp := 0
	for i := 1; i < l1+1; i++ {
		// last每次列下探时重置，记录左上角的数
		last := 0
		for j := 1; j < l2+1; j++ {
			// tmp表示正上方的数，即原dp[i-1][j]
			tmp = dp[j]
			// 当前遍历的字符相等，则取上一个公共值（未到当前lcs字符串之前的值）加1
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1
			} else {
				// 不相等，则等于相邻当前对比获得的最大的公共数
				dp[j] = max(tmp, dp[j-1])
			}
			last = tmp
		}
	}
	return dp[l2]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
