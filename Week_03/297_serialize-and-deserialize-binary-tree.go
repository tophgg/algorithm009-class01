package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法1：深度优先搜索 DFS
// 时间复杂度：O(N)，空间复杂度：O(N)。
// 解法2：迭代
type Codec struct {
	// deserialize(string) *TreeNode
	// serialize(*TreeNode) string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// 递归序列号
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// 分割数据，迭代数组
	vals := strings.Split(data, ",")
	return dfs(&vals)
}

func dfs(valsPtr *[]string) *TreeNode {
	// 第一个节点为根节点，剩余节点为左右节点
	val := (*valsPtr)[0]
	*valsPtr = (*valsPtr)[1:]
	// 遇到#终止， #标识空节点
	if val == "#" {
		return nil
	}
	// 获取第一个节点的值
	valInt, _ := strconv.Atoi(val)
	node := &TreeNode{Val: valInt}
	node.Left = dfs(valsPtr)
	node.Right = dfs(valsPtr)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */