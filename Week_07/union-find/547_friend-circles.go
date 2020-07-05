package main

// 并查集 未优化-时间复杂度O(N^3) 空间复杂度O(N)，优化路径压缩后 时间复杂度O(N)
// 纯dfs 时间复杂度O(N^2) 空间复杂度O(N)
var Parent = []int{}
var Count = 0
func findCircleNum(M [][]int) int {
	Count = len(M)
	Parent = make([]int, Count)
	Init(Count)
	// 初始化默认n个人 n圈子
	ans := Count
	// 对称二维矩阵1，2互为朋友为1  和2，1互为朋友也是1
	for i:=0; i<Count; i++ {
		for j:=i; j<Count; j++ {
			// 如果互为朋友，则将他们联通
			if M[i][j] == 1 {
				Union(i, j, &ans)
			}
		}
	}
	return ans
}

func Init(count int) {
	for i:=0; i<count; i++ {
		Parent[i] = i
	}
}

func Find(p int) int {
	for p!= Parent[p] {
		// 压缩
		Parent[p] = Parent[Parent[p]]
		p = Parent[p]
	}
	return p
}

func Union(p, q int, ans *int) {
	rootP := Find(p)
	rootQ := Find(q)
	// 如果已经联通则跳过
	if rootP == rootQ {
		return
	}
	// 发现一个非联通但是是朋友关系，则圈子-1
	Parent[rootP] = rootQ
	*ans -= 1
}

