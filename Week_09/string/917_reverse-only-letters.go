package main

import "unicode"

// 双指针 O(N),O(N)
// 栈顶记录，然后倒序逐个出栈
func reverseOnlyLetters(S string) string {
	l, r := 0, len(S)-1
	bytes := []rune(S)
	for l<r {
		if !unicode.IsLetter(bytes[l]) {
			l++
			continue
		}
		if !unicode.IsLetter(bytes[r]) {
			r--
			continue
		}
		if unicode.IsLetter(bytes[l]) && unicode.IsLetter(bytes[r]) {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
		}
	}
	return string(bytes)
}