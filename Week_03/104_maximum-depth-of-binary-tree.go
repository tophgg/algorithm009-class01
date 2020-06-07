package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法1：递归
// 时间复杂度：O(N)，空间复杂度：O(N)，最好情况平衡O(log(N))。
// 解法2：迭代
// 时间复杂度：O(N)，空间复杂度：O(N)。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {return a}
	return b
}