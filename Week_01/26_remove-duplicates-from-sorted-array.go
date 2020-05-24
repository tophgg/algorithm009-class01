package main

// 时间复杂度O(n) 空间复杂度O(1)
func removeDuplicates(nums []int) int {
	// j负责记录覆盖不重复的元素，从0开始； i从1开始与j位置的元素比较
	i, j, len := 1, 0, len(nums)
	for ; i<len; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j+1
}
