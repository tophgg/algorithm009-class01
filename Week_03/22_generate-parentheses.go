package main

var result []string
func generateParenthesis(n int) []string {
	result = []string{}
	_generate(0, 0, n, "")
	return result
}

func _generate(left int, right int, max int, str string) {
	if left == max && right == max {
		result = append(result, str)
	}

	if left < max {
		_generate(left+1, right, max, str + "(")
	}
	if left > right {
		_generate(left, right+1, max, str + ")")
	}
}
