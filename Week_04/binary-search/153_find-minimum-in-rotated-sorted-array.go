package main

// 解法1：二分法，找右侧的最小值
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		// 比较最右侧的数是否比中位数大，是的话  说明当前当前有序，右侧移动
		// 中位数比右侧大，说明中位数在左侧有序中，右移
		if nums[r] < nums[mid] {
			l = mid + 1
		} else {
			r--
		}
	}
	return nums[l]
}