package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法1：递归（深度优先遍历），每个元素都必须访问一次，最坏情况O(h)为树高，树为链表
// 时间复杂度：O(N),空间复杂度：O(N)
// 解法2：迭代，广度优先，使用栈，层层扫荡，先将根节点放入队列，不断迭代队列中的元素，对当前元素调换其左右字数的位置，之后若左子树不为空则放入队列。。。
// 时间复杂度：O(N),空间复杂度：O(N)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// right := invertTree(root.Right)
	// left := invertTree(root.Left)
	// root.Left = right
	// root.Right = left

	// 交换当前的左右子节点
	root.Right, root.Left = root.Left, root.Right
	// 下沉，交换左子树
	invertTree(root.Left)
	// 下沉，交换右子树
	invertTree(root.Right)
	return root
}


// 迭代
//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	stack := []*TreeNode{root}
//	for len(stack) > 0 {
//		node := stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		node.Left, node.Right = node.Right, node.Left
//		if node.Left != nil {
//			stack = append(stack, node.Left)
//		}
//		if node.Right != nil {
//			stack = append(stack, node.Right)
//		}
//	}
//	return root
//}