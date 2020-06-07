package main
// 子集
// 解法1：栈
// 时间复杂度：O(N+2^N)，空间复杂度:O(N+2^N))
// 解法2：递归、回溯
// 时间复杂度：O(N+2^N) 生成所有子集，并复制到输出集合中。空间复杂度:O(N+2^N))，存储所有子集，共 nn 个元素，每个元素都有可能存在或者不存在。
// 解法3：二进制位运算
// 时间复杂度：O(N+2^N)，空间复杂度:O(N+2^N))

func subsets(nums []int) [][]int {
	result := [][]int{}
	dfs(nums, 0, &result, []int{})
	return result
}

func dfs(nums []int, level int, result *[][]int, tmp []int) {
	subRes := make([]int, len(tmp))
	copy(subRes, tmp)
	// 每次递归都记录当前层结果的拷贝，包括第一次的[]结果
	*result = append(*result, subRes)

	// 每次递归从当前层开始，每次进来i都不同，所以不会存在重复记录
	for i:=level; i<len(nums); i++ {
		// 记录当前结果
		tmp = append(tmp, nums[i])
		// 继续下一步
		dfs(nums, i+1, result, tmp)
		// 恢复初始状态
		tmp = tmp[:len(tmp)-1]
	}
}