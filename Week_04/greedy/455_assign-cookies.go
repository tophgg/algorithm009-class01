package main

// 解法1：贪心+dfs 尽量多的填饱小朋友的肚子
// 解法2：dp
func findContentChildren(g []int, s []int) int {

	sort.Ints(s)
	sort.Ints(g)
	gi, si := 0, 0
	// 胃口小于饼干，喂！
	for len(g) > gi && len(s) > si {
		if g[gi] <= s[si] {
			gi++
		}
		si++
	}
	return gi
}