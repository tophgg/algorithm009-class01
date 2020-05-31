package main

// 解法1：递归
// 解法2：迭代
//
// 后序，左右根
//var res []int
//func postorder(root *Node) []int {
//	res = []int{}
//	dfs(root)
//	return res
//}
//
//func dfs(root *Node) {
//	if root != nil {
//		for _,n := range root.Children {
//			dfs(n)
//		}
//		res = append(res,root.Val)
//
//	}
//}


// 后序，左右根
//func postorder(root *Node) []int {
//	res := []int{}
//	if root == nil {
//		return res
//	}
//	for _, n := range root.Children {
//		res = append(res, postorder(n)...)
//	}
//	res = append(res, root.Val)
//	return res
//}


// 后序，左右根
func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res[] int
	var stack = []*Node{root}

	for 0 < len(stack)  {
		// 记录栈尾
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]
		//res = append([]int{node.Val}, res...)
		res = append(res, node.Val)

		for _, n := range node.Children {
			stack = append(stack, n)
		}
	}

	lp := 0
	rp := len(res) - 1
	for lp < rp { //指针相撞
		res[lp], res[rp] = res[rp], res[lp]
		lp++
		rp--
	}

	return res
}