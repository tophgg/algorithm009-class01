package main

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// 思路：p，q在同一边则其中一位是他们的公共祖先，pq不在同一边说明根节点就是公共祖先
// 1.递归，如果当前节点等于p或q则返回当前节点
// 时间复杂度：O(N),空间复杂度度O(N)，二叉树的节点数，最坏情况是树是一条链
// 2.存储父节点，p，q分别向上找父节点并记录map
// 时间复杂度：O(N),空间复杂度度O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 找不到，或者找到了（左节点或右节点），返回当前的值
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 左子树查找
	lson := lowestCommonAncestor(root.Left, p, q)
	// 右子树查找
	rson := lowestCommonAncestor(root.Right, p, q)

	// 向上返回，如果左右为nil了  说明都在剩下的右边那侧的结果
	if lson == nil {
		return rson
	}
	if rson == nil {
		return lson
	}
	// 左右都找到了，返回当前值，这个是第一次找到公共祖先时返回
	return root
}


// 存储父节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	// 递归存入parent为root的所有子节点
	dfs = func(r *TreeNode) {
		if r == nil {return}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)
	// visited记录p所有存放的值
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	// 遍历q如果存在在p的记录，则返回该节点
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}



