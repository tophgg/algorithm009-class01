package main


func climbStairs(n int) int {
	return climb(n)
}

func climb(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return climb(n-1) + climb(n-2)
}

func climbStairs2(n int) int {

	if n <= 2 {
		return n
	}

	first, second, third := 1, 2, 3


	for i := 3; i<=n; i++ {
		third = first + second
		first = second
		second = third
	}
	return third
}
