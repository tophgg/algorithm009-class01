package main

// 解法1：二分，思路：以target为目标点找出对半分，查看目标点与中点的比较，如果目标比中点小且比最左侧小说明不在右侧，直接右侧重复查找；如果比中点大，左侧有序，继续从右侧查找，如果比最右侧都大  可以提前结束。
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		}
		// 小于等于，中点比最左侧大，判断当前部分左~中是否有序
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 比最左侧还小，那说明mid在右侧，已经旋转。
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
