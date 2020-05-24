package main
// 解法1：先合并再排序 时间复杂度：O((N+M)log(N+M)) 空间复杂度：O(1)
// 解法2：双指针 从前往后，需要存放m个空间 时间复杂度：O(N+M) 空间复杂度：O(M)
// 解法2：双指针 从后往前，nums1中原地操作插入，注意nums2未插入完毕，需要循环继续插入 O(N+M) 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int)  {
	n1 := m-1
	n2 := n-1
	l := m+n-1
	// 注意0临界点，因为n1，n2已经-1了
	// 判断n1，n2把大的节点塞到num1的最后，并且两个长度-1
	for n1>=0 && n2>=0 {
		if nums1[n1] >= nums2[n2] {
			nums1[l] = nums1[n1]
			n1--
		} else {
			nums1[l] = nums2[n2]
			n2--
		}
		l--
	}
	// 可能出现nums的元素会比nums1小， 这时候
	for n2 >= 0 {
		nums1[l] = nums2[n2]
		n2--
		l--
	}
}
