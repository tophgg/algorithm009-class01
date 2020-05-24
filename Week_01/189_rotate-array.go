package main

func rotate(nums []int, k int)  {
	// 解法1， 反转 时间复杂度O(n) 空间复杂度O(1))
	// 解法2， 暴力-交换，超出范围要限制长度，移动i遍历，每次将i与要旋转的值交换，要交换k次，不断后移
	// 解法3， 环状替换，

	l := len(nums)
	k %= l // 防止k超出l的大小
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}