package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 原地翻转，迭代法， 时间复杂度O（n）, 空间复杂度O(1)
// 翻转口诀：斩断后路，不忘前事，才能重获新生
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for (head != nil) {
		// 保存前进方向， 保存下一跳，新的头部
		newHead := head.Next
		// 把当前指向上一个的位置，倒序
		head.Next = pre
		// 更新前驱指针
		pre = head
		// 当前指针，继续前进
		head = newHead
	}
	return pre
}


// 递归 时间复杂度O（n）空间复杂度O（n）
//func reverseList(head *ListNode) *ListNode {
//	// 标识结束递归，到最后一个节点
//	if head == nil || head.Next == nil {
//		return head
//	}
//	// cur是不变的，到最后一个点之后，结束递归，每层递归cur都是5，而head是不断在向前翻转，最终返回的cur就是翻转的链表
//	// 递归n次，cur都是最后一个节点开始
//	cur := reverseList(head.Next)
//	// 把下一个节点指向当前节点
//	head.Next.Next = head
//	// 防止链表循环，把当前节点设置为空
//	head.Next = nil
//	// 每层递归函数都返回cur，最后一个节点
//	return cur
//}