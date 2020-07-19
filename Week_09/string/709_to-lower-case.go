package main

func toLowerCase(str string) string {
	bytes := []rune(str)
	for i:=0; i<len(bytes); i++ {
		// 小写在大写后32位
		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] += 32
		}
	}
	return string(bytes)
}