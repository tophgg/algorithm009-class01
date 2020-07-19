package main

import (
	"math"
	"strings"
)

// 时间复杂度O(N)
func myAtoi(str string) int {
	return convert(clean(str))
}

func clean(s string) (sign int, abs string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	// 判断前置的符号位或数字？
	switch(s[0]) {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		abs=""
		return
	}
	for i, b := range abs {
		// 遍历到第i个有效的位置，就取s[:i]，剩下的为无效数字，可能为空格或字符串
		if b < '0' || '9' < b {
			abs = abs[:i]
			break
		}
	}
	return
}

// 接受处理有效的纯数字，判断是否超出限制
func convert(sign int, absStr string) int {
	absNum := 0
	for _, b := range absStr {
		// 累加构造10进制字符串的数
		absNum = absNum*10 + int(b-'0')
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}