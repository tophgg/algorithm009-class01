package main

// 逆序对，可用归并排序，每次合并前继续 nums[i] > 2 * nums[j]的个数
func reversePairs(nums []int) int {
	left, right := 0, len(nums)-1
	return mergeSortCount(nums, left, right)
}

func mergeSortCount(nums []int, left int, right int) int {
	if left >= right {return 0}
	mid := (left+right) >> 1
	count := mergeSortCount(nums, left, mid) + mergeSortCount(nums, mid+1, right)
	i:=left
	j:=mid+1
	for i<=mid && j<=right {
		// 因为左右两侧都是已排序好的数组，i++到下一个满足逆序对的节点
		if nums[i] <= 2*nums[j] {
			i++
		} else { // 因为i~mid为升序，所以i~mid即mid-i+1都符合 nums[i] > 2nums[j] 所以都加入计数中
			count += (mid-i+1)
			j++
		}
	}
	merge(nums, left, mid, right)
	return count
}

func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i:=left
	j:=mid+1
	k:=0
	for i<=mid && j<=right {
		if nums[i] > nums[j] {
			tmp[k] = nums[j]
			k++; j++
		} else {
			tmp[k] = nums[i]
			k++; i++
		}
	}
	for i<=mid {
		tmp[k] = nums[i]
		k++; i++
	}
	for j<=right{
		tmp[k] = nums[j]
		k++; j++
	}
	for p:=0; p<len(tmp); p++ {
		nums[left+p] = tmp[p]
	}
}