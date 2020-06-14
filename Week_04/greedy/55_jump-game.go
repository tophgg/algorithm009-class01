package main

// 解法1：贪心dfs，遍历存储  找出x + nums[x] >= nums，最远能到达某个位置，就能到达它前面的任何一个位置
// 时间复杂度：O(N)，空间复杂度：O(1)
// 解法2：dp
func canJump(nums []int) bool {
	maxIndex := 0
	lastIndex := len(nums) - 1
	for i, v := range nums {
		// 剪枝，无法到达下个目标点
		if i > maxIndex {
			return false
		}
		// 当前位置+当前可跳最远距离，比较历史最远位置，进行更新
		currIndex := i + v
		if maxIndex < currIndex {
			maxIndex = currIndex
		}
		// 剪枝，已经超过目标距离
		if maxIndex >= lastIndex {
			return true
		}
	}
	return true
}
