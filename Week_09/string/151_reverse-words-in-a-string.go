package main

import "strings"

// 1.手动deque，2.api
// 时间复杂度O(N)，空间复杂度O(N)
func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	l:=0
	r:=len(s)-1
	for l<=r && s[l] == ' ' {
		l++
	}
	for l<=r && s[r] == ' ' {
		r--
	}
	var deque []string
	word := ""
	for l<= r {
		c := s[l]
		if len(word)!=0  && c==' ' {
			deque = append([]string{word}, deque...)
			word = ""
		} else if c!=' ' {
			word += string(c)
		}
		l++
	}
	deque = append([]string{word}, deque...)
	return strings.Join(deque, " ")
}

// 使用函数
func reverseWords(s string) string {
	strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	letter := make([]string, 0)
	for i:=len(arr)-1; i>=0; i-- {
		if arr[i] == "" {
			continue
		}
		letter = append(letter, arr[i])
	}
	return strings.Join(letter, " ")
}
