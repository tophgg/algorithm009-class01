package main

// 解法1：二分法，判断每行第一个值,当left=right的时候  就可以在当前层二分继续查找
// 比如 matrix为m列=3， n行=4，mid的x=5/4=1， y=5%4=1
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	left, right := 0, m * n - 1
	for left <= right {
		mid := left + (right - left) / 2
		x := mid / n
		y := mid % n
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target{
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
