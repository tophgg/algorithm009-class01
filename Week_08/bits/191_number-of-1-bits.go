package main

// 解法1：掩码与运算，判断每个位置是否为1，是则有值； 时间复杂度O(1),空间复杂度O(1)
// 解法2：遍历判断是否1
// 递归
// 位运算不断去除最后一个数
func hammingWeight(num uint32) int {
	bits := 0
	var mask uint32 = 1
	for i:=0; i<32; i++ {
		if (num & mask) != 0 {
			bits++
		}
		mask <<= 1
	}
	return bits
}

// 位运算不断去除最后一个数
func hammingWeight(num uint32) int {
	bits := 0
	for num != 0 {
		bits++
		num &= (num-1) // n & n-1会将最后一位最低位去掉，比如010100 & 010011 => 010000 最后只要判断num是否变为0即可
	}
	return bits
}