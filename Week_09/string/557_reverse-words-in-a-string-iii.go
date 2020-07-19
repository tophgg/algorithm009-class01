package main
// O(N),O(N)
func reverseWords(s string) string {
	word := ""
	res := ""
	for i:=0; i<len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else {
			res += reverse(word) + " "
			word = ""
		}
	}
	res += reverse(word)
	return res
}

func reverse(word string) string {
	l:=0
	r:=len(word)-1
	bytes := []rune(word)
	for l<r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	return string(bytes)
}
