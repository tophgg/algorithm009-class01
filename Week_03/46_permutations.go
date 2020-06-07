package main

// 全排列
// 解法1：回溯，交换最左和最优
// 时间复杂度：O(N*N!)，空间复杂度：O(N）
// 解法2：用额外的切片来记录是否访问过对应的节点，临时栈空间
// 时间复杂度：O(N*N!) + O(N)，空间复杂度：O(N）
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	tmp := make([]int, len(nums))
	backstrack(0, nums, &result, tmp)
	return result
}

func backstrack(level int, nums []int, result *[][]int, tmp []int) {
	// 如果足够个数了，将临时tmp加入到result中
	if len(tmp) == level {
		sub := make([]int, level)
		copy(sub, tmp)
		*result = append(*result, sub)
		return
	}
	// level表示当前层的个数
	for i:=level; i<len(nums); i++ {
		// 每个循环
		tmp[level] = nums[i]
		// 交换本层和i的位置
		nums[level], nums[i] = nums[i], nums[level]
		backstrack(level+1, nums, result, tmp)
		nums[level], nums[i] = nums[i], nums[level]
	}
}


// 使用额外的数组记录是否使用过
func permute(nums []int) [][]int {
	used := make([]int, len(nums))
	result := make([][]int, 0)
	curNums := make([]int, 0)
	backstrack(0, nums, used, &result, curNums)
	return result
}

func backstrack(level int, nums []int, used []int, result *[][]int, curNums []int) {
	if len(curNums) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, curNums)
		*result = append(*result, tmp)
		return
	}
	// 从0~n
	for i:=0; i<len(nums); i++ {
		// 如果未用过才处理
		if used[i]== 0 {
			// 加入当前数字到记录里，设置已使用
			curNums = append(curNums, nums[i])
			used[i]=1
			// 进入下一层
			backstrack(level+1, nums, used, result, curNums)
			// 清空当前层
			used[i]=0
			curNums = curNums[:len(curNums)-1]
		}
	}
}