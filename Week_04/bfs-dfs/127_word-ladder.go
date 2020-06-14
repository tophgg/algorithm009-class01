package main

// BFS Time Complexity: O(n*26^l), l = len(word), n=|wordList| Space Complexity: O(n)
// 普通广度优先搜索依赖每层节点的分支数量。如果节点分支数量相同，搜索空间会随着层数的增加指数级的增加
// Bidirectional BFS Time Complexity: O(n*26^l/2), l = len(word), n=|wordList| Space Complexity: O(n)
// 解法1：双向bfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool)
	for _, word := range wordList {
		dict[word] = true
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	q1[beginWord] = true
	q2[endWord] = true
	l := len(beginWord)
	steps := 0

	for len(q1) > 0 && len(q2) > 0 {
		steps++
		if len(q1) > len(q2) { // 交换 分别走头开始走、下次从尾部开始走
			q1, q2 = q2, q1
		}
		q := make(map[string]bool) // 临时
		for k := range q1 {
			chs := []rune(k)
			for i:=0; i<l; i++ {
				ch := chs[i]
				for c := 'a'; c<='z'; c++ {
					chs[i] = c // 替换为新单词
					t := string(chs)
					if _,ok := q2[t]; ok { // 如果当前q1构造的单词，在q2中出现过，则说明已经相遇
						return steps+1
					}
					if _,ok := dict[t]; !ok { // 不存在字典的话 替换下一个字母
						continue
					}
					delete(dict, t) // 字典存在，删除该单词 -- 不允许重复
					q[t] = true // 单词加入临时集合中，
				}
				chs[i] = ch // 新单词第i位复位，继续替换下一个单词
			}
		}
		q1 = q // 将q1修改为新拓展的q（存放一系列符合）
	}
	return 0
}