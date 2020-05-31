package heap

import "container/heap"

// 解法1：小顶（根）堆，统计每个元素的出现频率，初始化小顶堆，存入k个元素，更新堆，每放入一个新元素，和堆顶比较，如果比堆顶大则替换堆顶元素，然后向下进行堆化，最终这个小顶堆就是前k高频的元素
// 时间复杂度O(Nlog(K))，实际为O(N+Nlog(K))，空间复杂度O(N))
// 解法2：快排变形 O(N)~O(Nlog(N))
// 解法3：map记录，然后倒排 时间复杂度O(N+NLOG(N))
func topKFrequent(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return make([]int, 0)
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	// 最小堆构建
	h := &NodeHeap{}
	topK := min(k, len(m))
	size := 0
	for k, v := range m {
		if size < topK {
			heap.Push(h, &Node{
				val: k,
				times: v,
			})
			size++
		} else {
			if v > (*h)[0].times {
				heap.Pop(h)
				heap.Push(h, &Node{
					val: k,
					times: v,
				})
			}
		}
	}
	res := make([]int, 0, topK)
	for i:=0; i<topK; i++ {
		res = append(res, heap.Pop(h).(*Node).val)
	}
	return res
}

type Node struct {
	val int
	times int
}

type NodeHeap []*Node

func (h NodeHeap)Len() int {
	return len(h)
}

func (h NodeHeap)Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h NodeHeap)Less(i, j int) bool {
	return h[i].times < h[j].times
}

func (h *NodeHeap)Push(x interface{}) {
	*h = append(*h, x.(*Node)) // 判断是否为Node类型
}

func (h *NodeHeap)Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(a, b int) int {
	if a < b {return a}
	return b
}
