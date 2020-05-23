package main

func MoveZeroes(nums []int) {
	if len(nums) <= 1 {
		log.Print(nums)
		return
	}
	count := 0
	for i, _ := range nums {
		// count记录0的个数，i-count为第一个0的地方，每遇到一个非0，且0个数>1的数就去跟最前一个0交换位置
		if nums[i] == 0 {
			count++
		} else if count > 0 { // nums[i] != 0 且有多余的0，交换当前位置和第一个0的位置（i-count）
			nums[i-count], nums[i] = nums[i], 0
		}
	}
	return
}