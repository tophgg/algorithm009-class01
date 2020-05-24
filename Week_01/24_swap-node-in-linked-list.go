package main

//func swapPairs(head *ListNode) *ListNode {
//	var prev *ListNode = &ListNode{0, head}
//	var hint *ListNode = prev
//	for prev.Next != nil && prev.Next.Next != nil {
//		prev.Next.Next.Next, prev.Next.Next, prev.Next, prev  = prev.Next, prev.Next.Next.Next, prev.Next.Next, prev.Next
//	}
//	return hint.Next
//}



// 前置指针迭代
// 原地翻转，迭代法， 时间复杂度O（n）, 空间复杂度O(1)
// 迭代
//func swapPairs(head *ListNode) *ListNode {
//	var dummy = &ListNode{}
//	dummy.Next = head
//	// 记录前驱节点
//	prev_node := dummy
//	for head != nil && head.Next != nil {
//		first_node := head
//		second_node := head.Next
//		// 交换
//		prev_node.Next = second_node
//		first_node.Next = second_node.Next
//		second_node.Next = first_node
//
//		// 更新前驱节点到交换后的头
//		prev_node = first_node
//		head = first_node.Next
//	}
//	return dummy.Next
//}

// 递归
// 递归法， 时间复杂度O（n）, 空间复杂度O(n)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}