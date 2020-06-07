package main

// 解法1：回溯,剪枝，此题与子集类似，但是存在k值相等时退出，还有剪枝优化问题，当i走到n-(k-len(tmp))<i时，后续的结果不用判断 max(i) = n - (k - pre.size()) + 1
// 时间复杂度： O(K*CN^K)
// 解法2：二进制
// 解法3：广度优先
func combine(n int, k int) [][]int {
	var result [][]int
	var tmp []int
	if n <= 0 || k <= 0 || n < k {
		return result
	}
	backstrack(1, n, k, &result, tmp)
	return result
}

func backstrack(level int, n int, k int, result *[][]int, tmp []int) {
	if len(tmp) == k {
		sub := make([]int, k)
		copy(sub, tmp)
		*result = append(*result, sub)
		return
	}
	for i:=level; i<=n-(k-len(tmp))+1; i++ {
		tmp = append(tmp, i)
		backstrack(i+1, n, k, result, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}