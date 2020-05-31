package main

// 解法1：递归
// 解法2：迭代（stack），把如果根节点不为空，左节点入栈，然后输出根节点，右节点入栈，并把当前根节点出栈
// 解法3：Morris 破坏树结构 （二叉树 转 单向升序链表）
// 解法4：Morris 保持树结构

// 解法1：递归
// 时间复杂度O(N)，空间复杂度O(logN) ~ O(N)
//func inorderTraversal(root *TreeNode) []int {
//
//	res := []int{}
//	helper(root, &res)
//	return res
//}
//
//func helper(root *TreeNode, res *[]int) {
//	if root == nil {
//		return
//	}
//	helper(root.Left, res)
//	*res = append(*res, root.Val)
//	helper(root.Right, res)
//}



// 解法1：递归
// 时间复杂度O(N)，空间复杂度O(logN) ~ O(N)
// 解法2：迭代（stack），把如果根节点不为空，左节点入栈，然后输出根节点，右节点入栈，并把当前根节点出栈
// 时间复杂度O(N)，空间复杂度O(N)
// 解法3：Morris 破坏树结构 （二叉树 转 单向升序链表）
// 时间复杂度O(N)，空间复杂度O(logN) ~ O(N)
// 解法4：Morris 保持树结构
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	// 首次stack无元素，需要root不为空作为判断
	for 0 < len(stack) || root != nil {
		// 先把左边的节点都压入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 当节点为空（左节点），这时候不断处理栈中元素，写入当前节点值，把节点指向栈顶元素的右节点
		// 保证每次都先处理左~中~右的顺序
		index := len(stack)-1
		res = append(res, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
	}
	return res
}
