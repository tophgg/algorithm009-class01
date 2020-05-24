package main

// 时间复杂度O(n) 空间复杂度O(1)
func plusOne(digits []int) []int {
	// 等于0时，要走+1的逻辑，所以循环条件为>=0
	for i:=len(digits)-1; i>=0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}