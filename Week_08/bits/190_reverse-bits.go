package main

// 解法1：取模求和
// 解法2：按位翻转 时间复杂度O(logN),空间复杂度O(1)
// 解法3：分治1~2~4~16不断交换两侧的值
func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		// 从右往左开始累加，将当前最右边的位值 右移power位到对应的地方
		ret += (num & 1) << power
		// 移动后，num右移，power浮标也右移  相当于双指针
		num >>= 1
		power -= 1
	}
	return ret
}