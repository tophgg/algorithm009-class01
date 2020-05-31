package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
解法1：递归
解法2：迭代
 */
//var res []int
//func preorderTraversal(root *TreeNode) []int {
//	res = []int{}
//	helper(root)
//	return res
//}
//
//func helper(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	res = append(res,root.Val)
//	helper(root.Left)
//	helper(root.Right)
//}

// 迭代
// 前序遍历，先打印该节点，然后是它的左子树，最后才是右子树（入栈）
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		// 当前节点先输出，当前节点不为空，先把右节点都压入栈中，节点指向左
		// 保证每次先输出根节点，再处理左节点，然后是右节点
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}

		// 当左节点为空了，这时候处理栈中的元素root指向栈中最后的右节点，
		index := len(stack) - 1
		root = stack[index] // 栈顶出栈
		stack = stack[:index]
	}
	return res
}