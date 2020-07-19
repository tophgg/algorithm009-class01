package main

// 1.滑动窗口
// 2.hash存储
func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	need := make(map[byte]int)
	// 用hash记录需要的字符串
	for v := range p {
		need[p[v]]++
	}
	l, r := 0, 0
	valid := 0
	res := []int{}
	// 先比较右侧，加入窗口元素
	for r < len(s) {
		// c表示移入窗口的字符
		c := s[r]
		r++
		// 找到适合的，就移入窗口
		if _, ok := need[c]; ok {
			window[c]++
			// 如果当前字符的哈希表想同说明当前窗口找到一个有效的异位词
			if window[c] == need[c] {
				valid++
			}
		}
		// 再比较滑动窗口是否超长，超长需要去除左侧元素
		for r-l >= len(p) {
			// 如果p的总数和有效的总数一致，说明当前的l就是有效的起点
			if len(need) == valid {
				res = append(res, l)
			}
			// 当前窗口已满，如果当前的字符是合适的字符，将移出窗口，如果是有用的词需要将有效词数-1
			d := s[l]
			l++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}