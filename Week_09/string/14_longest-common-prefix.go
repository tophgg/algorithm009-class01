package main

// 暴力匹配第一个和后面每一个
// trie
// 不断两两比较相临的字符 时间复杂度O(MN),空间复杂度O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {return ""}
	// 每个字符串都和第一个字符串相比较；如果第j个字符串到达末尾 || 第j个字符和第一个字符到某个i不匹配则退出
	for i:=0; i<len(strs[0]); i++ {
		for j:=1; j<len(strs); j++ {
			if len(strs[j]) == i || strs[j][i] != strs[0][i] {
				return strs[j][:i]
			}
		}
	}
	return strs[0]
}

// 不断两两比较相临的字符 时间复杂度O(MN),空间复杂度O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i:=1; i<len(strs); i++ {
		prefix = Lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func Lcp(str1, str2 string) string {
	index := 0;
	length := min(len(str1), len(str2))
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(a, b int) int {
	if a>b {return b}
	return a
}