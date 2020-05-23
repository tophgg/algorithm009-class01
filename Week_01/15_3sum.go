package main

import "sort"

// 夹逼解法 时间复杂度O(n^2)  空间复杂度O(1)  实际时间复杂度为排序+数组*双指针遍历：O(nlogn) + O(n) * O(n)
// https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	ret := make([][]int, 0, 0)
	for i := 0; i<len(nums)-1; i++ {
		l, r := i+1, len(nums)-1
		// 过滤掉不可能的情况
		if nums[i] > 0 || nums[i] + nums[l] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
        // l，r双指针夹逼
		for l < r {
			// l去重
			if l > i + 1 && nums[l] == nums[l-1] {
				l++; continue
			}
			// r去重
			if r < len(nums) - 2 && nums[r] == nums[r+1] {
				r--; continue
			}

			// 都大于0不可能
			if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else if nums[i] + nums[l] + nums[r] < 0 {
				l++
			} else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				// 因为已经去重，所以l和r要同时放大，缩小才有可能加起来为0
				l++; r--
			}
		}
	}
	return ret
}

