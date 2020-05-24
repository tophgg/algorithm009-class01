package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1：迭代，用前驱指针记录链表起点，移动pre指针不断连接下一个较小的值
// 时间复杂度：O(n+m)  空间复杂度：O(1)
// 解法2：递归，找出递归出口，l1、l2为空，迭代递归的连接点l1.next为l1.next或l2，并返回l1（l2同理）
// 时间复杂度：O(n+m)  空间复杂度：O(n+m)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	preHead := pre
	// pre作为向后移动的标记游标，preHead用于记录链表起点
	for l1 != nil && l2 != nil {
		// l1，l2向前移动，pre指向较小的链表；pre移动
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	// 最后拼接剩余不为空的链表
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}
	return preHead.Next
}