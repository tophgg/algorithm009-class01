package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 问题：到叶子节点的最短距离
// 叶子节点的定义是左孩子和右孩子都为 null 时叫做叶子节点
// 当 root 节点左右孩子都为空时，返回 1
// 当 root 节点左右孩子有一个为空时，返回不为空的孩子节点的深度
// 当 root 节点左右孩子都不为空时，返回左右孩子较小深度的节点值
// 解法1：递归
// 时间复杂度:O(N)，空间复杂度O(logN)-平衡二叉树，不平衡时最差O(N)
// 解法2：广度优先搜索迭代
// 时间复杂度:O(N)，空间复杂度O(N)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dl := minDepth(root.Left)
	dr := minDepth(root.Right)
	if root.Left == nil {
		return dr + 1
	} else if root.Right == nil {
		return dl + 1
	} else {
		return min(dl, dr) + 1
	}
}

func min(a, b int) int {
	if a < b {return a}
	return b
}
