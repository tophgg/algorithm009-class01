package main

// 维护堆
func heapify(arr []int, length int, i int) {
	// 完整二叉堆，左右子树公式
	left := 2 * i + 1
	right := 2 * i + 2
	// 当前被交换的元素就是原最大堆的堆顶
	largest := i
	// 不断选举比较左右子树，找出更大的元素作为新堆顶
	if left < length && arr[left] > arr[largest] {
		largest = left
	}
	if right < length && arr[right] > arr[largest] {
		largest = right
	}
	// 如果新堆顶不等于当前i节点的位置，则替换最大堆节点，然后继续进行维护
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, length, largest)
	}
}

func heapSort(arr []int) {
	if len(arr) == 0 {return}
	length := len(arr)
	// 从底部构建堆，大顶堆；此时堆为无序区
	for i:=length/2 - 1; i>=0; i-- {
		heapify(arr, length, i)
	}
	// 从后往前，将最后一个元素置于堆顶，得到R1,R2...Rn-1为新堆；Rn为有序区
	// 不断调整前n-1的新堆，然后再将当前R1最大元素，和无序区最后一个元素调换，直到无序区遍历完成
	for i:=length-1; i>=0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}