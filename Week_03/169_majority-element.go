package main

// 解法1：哈希表，遍历两次
// 时间复杂度：O(N)，空间复杂度：O(N)
// 解法2：排序后，取n/2位置的元素为总数
// 时间复杂度：O(Nlog(N))，空间复杂度：O(Nlog(N))，堆排序的话空间为O(1)
// 解法3：随机化，随机挑选一个位置的元素进行验证是否为众数
// 时间复杂度：O(N),也可以趋向于O(+无穷)空间复杂度O(1)
// 解法4：分治，
// 时间复杂度:O(N)，空间复杂度O(logN)，用到递归栈空间
// 解法5：投票算法，利用众数出现n/2次，可得维护一个cor记录每个数字出现的count次，如果出现相同的数就+1，出现不同的数就-1，直到为0时换新的数进行累加，最后出来的数count会大于0，且这个数即为众数
// 时间复杂度：O(N),空间复杂度O(1)

func majorityElement(nums []int) int {
	major := 0
	counter := 0
	for i:=0; i<len(nums); i++ {

		if counter == 0 {
			major =  nums[i]
		}
		if major == nums[i] {
			counter++
		} else {
			counter--
		}
	}
	return major
}