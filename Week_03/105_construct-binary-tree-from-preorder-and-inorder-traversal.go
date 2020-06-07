package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思考：前序- 根左右，中序- 左根右
// 解法1：递归
// 时间复杂度：O(N)，空间复杂度：O(N)
// 解法2：迭代
// 时间复杂度：O(N)，空间复杂度：O(N)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 初始化第一个根节点为起点
	root := &TreeNode{preorder[0], nil, nil}
	// 找到中序根节点的位置
	r := 0
	for ; r<len(preorder); r++ {
		if inorder[r] == preorder[0] {
			break
		}
	}
	// 构造左右子树
	root.Left = buildTree(preorder[1:r+1], inorder[:r])
	root.Right = buildTree(preorder[r+1:], inorder[r+1:])
	return root
}