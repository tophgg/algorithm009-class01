package heap

// 解法1：使用栈（优先队列），维护一个最小k的栈，每次移动，加入元素并进行维护成最小堆，并将最小的堆顶pop出。适合数据量较大的数据不然直接排序会更快。堆仅占用额外k个数据的空间，而且不需要修改原有数据。
// 解法2：排序整个数组，把前k个数输出， 时间复杂度：O(NlogN) 空间复杂度O(1)需原地改变数组。
// 解法3：快排变形优化版，每次快排选择后，左边小于基准点V，右边大于基准点V，每次分割直到左侧加起来的数量=k。优化点，每次选用节点选择随机基准点。
// 时间复杂度O(N)，最坏为O(N^2)，如果每次划分都是最大值，要划分n-1次，则为O(N^2)最坏情况，一般情况为N+N/2+N/4+...N/N = 2N即为O(N)
// 空间复杂度O(logN)，期望深度，但是如果需要划分n次，则每层占用O(1)的空间需要O(N)

//利用快排思想，每次只处理一个部分
//补充：实际下标为k-1，但获取前k+1个也就获取了k个，-1操作反而消耗额外时间
//实际结果，将k换为k-1，时间消耗从2ms上升至3ms
//具体思路：
//当快排后的枢纽位置在k，恰好左边部分为要求的k+1个数
//当枢纽位置小于k，前几个数字已经确定，需要确定后几个数字，令low = pivot + 1继续
//当枢纽位置大于k，前k个数字在枢纽前，在枢纽前寻找数字，令high = pivot - 1继续
//下面讨论low < high作为循环结束条件的合理性：
//一般情况下，若k小于数组长度，low必然<=k，high必然>=k
//由于pivot == k时直接结束循环不考虑，因此只有可能出现low == high == k
//要使得上式成立，必然有两个状态满足pivot1==k-1和pivot2==k+1
//pivot1前的部分必然是小于k的，而pivot2后的部分必然是大于k的
//因此可以确定k是pivot1后数字中最小的一个，可以确定前k+1个数字已满足要求
//特殊情况下，k等于数组长度。此时pivot < k始终成立，也就是pivot!=k
//每趟处理后，必然有low = pivot + 1，也就是low必然会移动到满足low>=high
//由于k等于数组长度，因此无论如何变动，取出的数组一定满足要求
func getLeastNumbers(arr []int, k int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		pivot := partition(arr, left, right)
		if pivot > k {
			right = pivot-1
		}
		if pivot < k {
			left = pivot + 1
		}
		if pivot == k {break}
	}
	return arr[:k]
}

func partition(nums []int, left int, right int) int {
	pivot := nums[left]
	// 不断位移
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
			for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	// 出来时low和right相同，这时这个节点就是中间节点
	nums[left] = pivot
	return left
}

// 维护堆

//func getLeastNumbers(arr []int, k int) []int {
//	if k > len(arr) || k == 0 {
//		return nil
//	}
//	// 前k个数组构建大顶堆？
//	res := arr[:k]
//	n := len(res)-1
//	for i:=len(res)/2 - 1; i>=0; i-- {
//		build(res, i, n)
//	}
//	// 遍历k个元素，调整堆
//	for i:=k; i<len(arr); i++ {
//		if arr[i]<res[0] {
//			res[0] = arr[i]
//			build(res,0,k-1)
//		}
//	}
//	return res
//}
//
//func build(nums []int, i int, n int) {
//	for {
//		index := 2*i + 1
//		// index超过k则退出
//		if index > n || index < 0 {
//			break
//		}
//		// index的子节点大于当前index节点，那么index指向子节点
//		if j:=index+1; j<=n && nums[j]>nums[index] {
//			index = j
//		}
//		// 当前节点大于子节点退出
//		if nums[i] > nums[index] {
//			break
//		}
//		// 交换当前节点和子节点
//		nums[index], nums[i] = nums[i], nums[index]
//		// 指向子节点
//		i = index
//	}
//}



