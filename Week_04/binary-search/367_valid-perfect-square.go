package main

func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := left + (right - left)/2
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}



// 解法1：二分
// 解法2：牛顿迭代
//func isPerfectSquare(num int) bool {
//	if num < 2 {return true}
//	x := num / 2
//	for x * x > num {
//		x = (x + num/x)/2
//	}
//	if x * x == num {
//		return true
//	}
//	return false
//}