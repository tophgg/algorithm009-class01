package main

// 解法1：递归
// 解法2：迭代
// 思路，先将节点写入栈，每次出栈先写入节点，然后倒序压入子节点入栈，保证每次出栈都是最左侧节点
// 前序，根左右
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	res = append(res, root.Val)
	for _, n := range root.Children {
		res = append(res, preorder(n)...)
	}
	return res
}


// 前序，根左右
//func preorder(root *Node) []int {
//	if root == nil {
//		return []int{}
//	}
//	var stack = []*Node{root}
//	var res []int
//	for len(stack) > 0 {
//		// 取出栈顶，直接将节点写入结果
//		index := len(stack) - 1
//		node := stack[index]
//		res = append(res, node.Val)
//		stack = stack[:index]
//		// 倒序插入子节点，保证每次出栈都是最左侧的节点，保证了根左右
//		for i:=len(node.Children)-1; i>=0; i-- {
//			stack = append(stack, node.Children[i])
//		}
//	}
//	return res
//}