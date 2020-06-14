package main

func jump(nums []int) int {
	maxInt, end, lastInt, steps := 0, 0, len(nums)-1, 0
	for i:=0; i<lastInt; i++ {
		// 更新最远距离
		if nums[i]+i > maxInt {
			maxInt = nums[i]+i
		}
		// 剪枝
		if maxInt >= lastInt {
			steps++
			break;
		}
		// 下一步，更新边界
		if end == i {
			end = maxInt
			steps++
		}
	}
	return steps
}