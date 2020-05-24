package main

// 解法1：哈希表，记录所有 时间复杂度O（n）空间复杂度O（n）
// 解法2：快慢双指针，比较前后指针遍历，是否会出现相等情况 时间复杂度O（n）空间复杂度O（1）
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int, 0)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		} else {
			m[head] = 1
		}
		head = head.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var (
		slow *ListNode = head
		fast *ListNode = head.Next
	)
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}