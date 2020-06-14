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
// 解法1：队列bfs
func largestValues(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		maxValue := math.MinInt32
		l := len(queue)
		for i:=0; i<l; i++ {
			maxValue = max(maxValue, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, maxValue)
		queue = queue[l:]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
