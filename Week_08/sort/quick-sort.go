package main

// 快排
func quickSort(arr []int, left int, right int) {
	if right <= left {
		return
	}
	pivot := patition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

func patition(arr []int, left int, right int) int {
	// 以最右侧的节点作为pivot，最左侧作为计数起始点，只要计数点小于pivot节点的值，计数点就右移；
	// 直到最后一个不小于的节点，则和pivot交换，此时counter的位置则是中点；左侧都是小于pivot的值，
	pivot := right
	counter := left
	for i:=left; i<right; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}


