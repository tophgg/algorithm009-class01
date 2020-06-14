package main

// 注意此题限制非负整数，包括0
// 解法1：二分法，正数测是单调递增的，而且又边界为0 ~ nums[last]，可根据index获得。 可以用位运算
// 解法2：牛顿迭代 cur = (cur+x/cur)/2
func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid * mid == x {
			return mid
		}
		if mid * mid > x {
			r = mid - 1
		} else {
			ans = mid
			l = mid + 1
		}
	}
	return ans
}
//
//func mySqrt(x int) int {
//	var a int = x
//	for a * a > x {
//		a = (a + x/a)/2
//	}
//	return a
//}