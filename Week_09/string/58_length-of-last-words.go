package main

func lengthOfLastWord(s string) int {
	if len(s) == 0 {return 0}
	last := len(s)-1;
	length := 0
	for last > 0 && s[last] == ' ' {
		last--
	}
	for i:=last; i>=0; i-- {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {
			length++
		} else {
			return length
		}
	}
	return length
}