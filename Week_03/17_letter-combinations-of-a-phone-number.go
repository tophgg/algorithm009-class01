package main

// 解法1：广度优先 3层循环，第一层循环digist，第二层循环phone[digist[i]]，第三层循环给已有字符串数组里拼接上第二层的num
// 解法2：递归四件套
// 解法3：回溯，每次将digits截取第一位，找到alpha中对应的字符串，然后循环字母拼接到下一次递归里，直到digits为空（暴力穷举）
// 时间复杂度：O(3^n + 4^m)，n是对应3个字母的数字的情况，m是对应4个字母的数字的情况
// 空间复杂度：O(3^n + 4^m)
// 解法4：队列、栈来操作
var digitAlpha = map[string]string{
	"2" : "abc",
	"3" : "def",
	"4" : "ghi",
	"5" : "jkl",
	"6" : "mno",
	"7" : "pqrs",
	"8" : "tuv",
	"9" : "wxyz",
}

func letterCombinations(digits string) []string {
	// map  数字和字母的关系
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{}
	_generate(digits, "", &result)
	return result
}

func _generate(digits string, str string, result *[]string) {
	// 终结者
	if len(digits) == 0 {
		*result = append(*result, str)
		return
	}
	// logic

	// drill
	for _, ch := range digitAlpha[string(digits[0])] {
		_generate(digits[1:], str + string(ch), result)
	}
	// clear
}
