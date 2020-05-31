package main

func isAnagram(s string, t string) bool {
	// 解法1：把字符串放到数字里，然后排序比较，时间复杂度O(NLOGN), 空间复杂度O(1)取决于函数的空间占用
	// 解法2：把字符串放到哈希表里，进行比较，如果出现次数不一样或不存在对应字母，则不一致，可以用ASCII优化；
	// 解法3：含有unicode字符，只能用map来记录
	if len(s) != len(t) {
		return false
	}
	a := [26]int{}
	b := [26]int{}
	for i:=0; i<len(s); i++ {
		a[s[i] - 'a']++
		b[t[i] - 'a']++
	}

	return a == b
}