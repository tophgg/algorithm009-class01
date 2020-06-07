package main

// 解法1：分治递归 冥函数二分时，分为奇数和偶数时要分别处理 ，注意负数的时候要取倒数
// 时间复杂度：O(logN) 空间复杂度：O(logN)
// 解法2：迭代 不占用额外空间，利用二进制的位数相乘
// 时间复杂度：O(logN) 空间复杂度：O(1)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return dfs(x, n)
	}
	return 1.0/dfs(x,-n)
}

func dfs(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	next := dfs(x, n/2)
	if n%2 == 0 {
		return next * next
	} else {
		return next * next * x
	}
}