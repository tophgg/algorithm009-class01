package main
// 1.哈希表 存放每个字符串出现的次数，然后遍历s一次字符串如果等于1则返回
// 2.暴力法 时间复杂度O(N^2)
func firstUniqChar(s string) int {
	// hash := make(map[byte]int, len(s))
	hash := [26]int{}
	for i:=0; i<len(s); i++ {
		hash[s[i]-'a']++
	}
	for i:=0; i<len(s); i++ {
		if hash[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}