package main

// dp dp[i]表示s[i]时子串的总长度，dp[i] = s[i-dp[i-1]-1] == s[i] ?
// 时间复杂度O(N^2)，空间复杂度O(N^2)
// 中心扩展法
// 时间复杂度O(N^2)，空间复杂度O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {return ""}
	// 截取子串为止 是否为回文，如果是的话 则进行判断
	// dp[start][end] => 0 | 1
	dp := make([][]int, len(s))
	for i:=0; i<len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	res := ""
	// 由短到长，l表示当前子串的到末尾的长度
	for l:=1; l<=len(s); l++ {
		for start:=0; start<len(s); start++ {
			// 结束为止索引
			end := start+l-1
			if end >= len(s) {
				break
			}
			// 单字符为回文
			if l== 1 {
				dp[start][end]=1
			} else if l == 2 {
				// 双字符相同时回文
				if s[start] == s[end] {
					dp[start][end] = 1
				}
			} else {
				// 多字符如果两边相同，由去掉的部分决定
				if s[start] == s[end] {
					dp[start][end] = dp[start+1][end-1]
				}
			}
			// 1说明有，
			if dp[start][end] == 1 && l>len(res) {
				res = s[start:end+1]
			}
		}
	}
	return res
}

// 中心拓展法，每个点作为中心都向两边拓展，直到无法拓展为止
func longestPalindrome(s string) string {
	maxLen := 0
	maxLenStart := 0
	for i:=0; i<len(s); i++ {
		// 开始的中心节点，和拓展的长度
		// 分别从单字符、双字符向外拓展
		start1, len1 := maxlen(s, i, i)
		start2, len2 := maxlen(s, i, i+1)

		// 取较大的，来跟最大值进行比较
		start, theMax := start1, len1
		if len2 > len1 {
			start, theMax = start2, len2
		}
		if theMax > maxLen {
			maxLen = theMax
			maxLenStart = start
		}
	}
	return s[maxLenStart:maxLenStart+maxLen]
}

// 开始节点向左移，结束节点向右移，直到不相等，不为回文；这时候返回对应的长度
func maxlen(s string, start, end int) (int, int) {
	for start>=0 && end<len(s) && s[start]==s[end] {
		start--
		end++
	}
	// 起始的位置 索引start+1；
	// 长度 (end-1) - (start+1) + 1
	return start+1, end-start-1
}

func max(a, b int) int {
	if a>b {return a}
	return b
}