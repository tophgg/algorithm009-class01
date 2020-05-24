package linked_list

// 返回连接节点 即快指针的位置？
// 1.找出相遇节点
// 2.从头开始走直到相遇
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	// 快慢指针，找到第一次相遇点 a=起点到环入口的距离   b=环长度  s=nb, f=2nb   f = s + nb
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}
	// 相遇 fast回到起点，让快慢指针再次相遇   s = a + nb为环入口
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}
