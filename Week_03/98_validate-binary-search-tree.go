package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 左子树比根节点小，右子树比根节点大
// 解法1：递归，
// 时间复杂度：O(N)，空间复杂度：O(N)。
// 解法2：中序遍历，左中右，是升序的，保存上一个根节点，进行中序遍历，
// 比较下一个节点是否比上一个根节点小即可
// 时间复杂度：O(N)，空间复杂度：O(N)。
func isValidBST(root *TreeNode) bool {
	// 初始 左节点无限小，右节点无限大
	return helper(root, math.MinInt64,math.MaxInt64)
}

func helper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	// 如果当前节点的值比最小值还小，或者比最大值还大，则不符合
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	// 向下沉，左节点设置最大值为当前；右节点设置最小值为当前
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}


// 中序遍历
//func isValidBST(root *TreeNode) bool {
//	stack := []*TreeNode{}
//	inorder := math.MinInt64
//	for len(stack) > 0 || root != nil {
//		for root != nil {
//			stack = append(stack, root)
//			root = root.Left
//		}
//		root = stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		if root.Val <= inorder {
//			return false
//		}
//		inorder = root.Val
//		root = root.Right
//
//	}
//}