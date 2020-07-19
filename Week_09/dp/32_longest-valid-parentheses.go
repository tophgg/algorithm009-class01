package main

// dp[i]表示以s[i]结尾的子串能形成的最长括号的长度
// O(N)、O(N)
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	dp := make([]int, len(s))
	maxVal := 0
	for i := 1; i < len(s); i++ {
		// 如果当前为（，说明当前的子串为0，不用判断
		// 如果当前为），判断上一个括号的情况
		if s[i] == ')' {
			// 如果上一个括号为（，则可以闭合，当前为1时，有2个，当前大于1时有dp[i-2]+2
			if s[i-1] == '(' {
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
				// 如果上一个括号为），则需要判断前面的括号是否能够闭合，找到上一个括号有效括号数的再上一个位置  判断是否为(
				// 此时需要先判断这个位置是否有效，然后加上这个回环前面的有效括号数，同样需要判断前面的回环是否有效
				// dp[i] = dp[i-1]+2 是因为))找到了前一个的一对括号对
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > maxVal {
			maxVal = dp[i]
		}
	}
	return maxVal
}

// 栈 O(N)、O(N)
func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		// 左括号入栈，右括号出栈
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 弹出栈顶
			stack = stack[:len(stack)-1]
			// 如果）是最后一个元素的话，要把当前的入栈，否则计算最近括号匹配数
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				// 当前元素的索引与栈顶元素作差，获取最近的括号匹配数 -- 当前i与栈顶之差的长度
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
