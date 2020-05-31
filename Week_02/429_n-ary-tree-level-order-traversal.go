package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 解法1：bfs广度优先搜索
// 时间复杂度：O(n)。n 指的是节点的数量。空间复杂度：O(n)。
// 解法2：递归
// 时间复杂度：O(n)。n 指的是节点的数量 空间复杂度：正常情况 O(\log n)，最坏情况 O(n)。运行时在堆栈上的空间。

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*Node{root}
	var level int
	// 广度优先，以队列先进先出来标记根节点
	for 0 < len(queue) {
		counter := len(queue)
		res = append(res, []int{})
		// 当前层队列里的val逐个放入当前层的数组中
		for i:=0; i<counter; i++ {
			// if queue[i] != nil {
			res[level] = append(res[level], queue[i].Val)
			// 将子节点都放入队列
			for _, n := range(queue[i].Children) {
				queue = append(queue, n)
			}
			// }
		}
		// 清空当前层队列
		queue = queue[counter:]
		// 计算下一层
		level++
	}
	return res
}

// 递归
// 子节点循环时，使用递归同一层的level相同，下一层时，需要先分配二维数组
//var result [][]int
//func levelOrder(root *Node) [][]int {
//	result = [][]int{}
//	if root != nil {a
//		traverse_node(root, 0)
//	}
//	return result
//}
//
//func traverse_node(root *Node, level int) {
//	if len(result) == level {
//		result = append(result, []int{})
//	}
//	result[level] = append(result[level], root.Val)
//	for _, child := range root.Children {
//		traverse_node(child, level+1)
//	}
//}