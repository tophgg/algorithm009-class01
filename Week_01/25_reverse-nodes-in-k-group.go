package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1：栈，每次维护长度为k个栈，出栈翻转栈内节点，最终出栈时分别指向头尾节点进行连接
// 需判断剩下链表是否大于k个；已翻转部分需要与剩下链表连接
// 解法2：
// 时间复杂读为O(n)，空间复杂度为O（1）
// 分为翻转部分+待翻转+未翻转  尾插法
// 设置前驱指针+起始节点+结束节点，局部翻转，然后再判后续k个是否为空，不空则继续设置节点，进行局部翻转
// 解法3：递归
//
func reverseKGroupTailInsert(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head

	var pre, end *ListNode = dummy, dummy
	// dummy(pre) -> head(start) -> .... -> end -> end.Next
	for end.Next != nil {
		// 后移k位到tail，当前最后翻转节点
		for i:=0; i<k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		// 将start-end，k个链表节点翻转，
		pre.Next = reverse(start)
		// 开始节点到最后，指向原来的标记的end.next
		start.Next = next
		// 初始前驱节点指向跳到最后翻转后的start位置
		pre = start
		// 初始end节点指向start位置，下次循环再移位
		end = pre
	}
	return dummy.Next
}

// 翻转辅助函数，pre -> cur -> cur.next    ===>  cur -> pre	===>   pre=cur, cur=next	继续判断下2个节点，最终pre为倒序
//
func reverse(head *ListNode) *ListNode{
	var pre *ListNode = &ListNode{}
	var curr *ListNode = head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}



func reverseKGroupStack(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for {
		count := k
		stack := []*ListNode{}
		tmp := head
		// 当链表足够长，入栈，链表后移
		for count > 0 && tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Next
			count -= 1
		}
		// 目前tmp所在k+1位置

		// 如果结束了，count还有，说明剩余不够k个位置
		if count > 0 {
			// 将p回到原来的节点
			// p.Next = head
			break
		}
		// 出栈，翻转操作
		// p -> next           <--->            p  ->  next
		// 3     2     1                 3 ->   2       1
		for i:=len(stack)-1; i>=0; i-- {
			p.Next = stack[i]
			p = p.Next
		}
		// 此时p为当前翻转的最后一个节点（下一次翻转的前驱节点），tmp为翻转节点下一个节点
		p.Next = tmp
		// 下一轮循环的头节点从tmp开始
		head = tmp
	}
	return dummy.Next
}