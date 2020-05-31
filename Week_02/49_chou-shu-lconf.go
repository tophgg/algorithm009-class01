package heap

import (
	"container/heap"
	"math"
)

// 解法1：最小堆，丑树的因数只有2，3，5，第一个丑数1*2，3，5可得到一批丑数，在这里选一个最小的数字作为下一个丑数，然后再乘以2，3，5可以得到一批新的丑数
// 可以用一个最小堆来维护这些数字集合，直到第n次循环来找到第n个丑数，时间复杂度为O(nlogn))，辅助最小堆，空间复杂度为O(N)
// 解法2：动态规划
// 时间复杂度为O(n)，空间复杂度为O(n)
func nthUglyNumber(n int) int {
	uglyNum := -1
	var h IntHeap
	heap.Init(&h)
	heap.Push(&h, 1) // 初始化1
	for n > 0 {
		for h[0] == uglyNum {
			// 如果堆顶等于上一个丑数，已有重复的值，就丢弃
			heap.Pop(&h)
		}
		uglyNum = heap.Pop(&h).(int)

		// 约束
		if 2*uglyNum <= math.MaxInt32 {heap.Push(&h, 2*uglyNum)}
		if 3*uglyNum <= math.MaxInt32 {heap.Push(&h, 3*uglyNum)}
		if 5*uglyNum <= math.MaxInt32 {heap.Push(&h, 5*uglyNum)}
		n-- // 计算出一个新的丑数后，n-1
	}
	return uglyNum
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) } // 入堆
func (h *IntHeap) Pop() interface{} { // 出堆
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}


// 解法2：动态规划
//func nthUglyNumber(n int) int {
//	dp := make([]int, n)
//	dp[0] = 1
//	idx2, idx3, idx5 := 0, 0, 0
//	for i:=1; i<n; i++ {
//		dp[i] = min(dp[idx2]*2, min(dp[idx3]*3, dp[idx5]*5))
//		if dp[i] == dp[idx2]*2 {idx2++}
//		if dp[i] == dp[idx3]*3 {idx3++}
//		if dp[i] == dp[idx5]*5 {idx5++}
//	}
//	return dp[n-1]
//}
//
//func min(a,b int) int {
//	if a < b {return a}
//	return b
//}