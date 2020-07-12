package main

func bubbleSort(arr []int) {
	len := len(arr)
	for i:=0; i<len-1; i++ {
		for j:=0; j<len-1-i; j++ { // -i是因为前面遍历后，都是最大的元素再后i个位置，无需重复遍历
			if arr[j] > arr[j+1] { // 比较相连元素，如果前面的比后面的大则交换；最终最大的元素都排在最后面
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func selectSort(arr []int) {
	len := len(arr)
	minIndex := 0
	for i:=0; i<len-1; i++ { // 前0~i保证存储最小的数
		minIndex = i
		for j:= i+1; j<len; j++ {
			if arr[j] < arr[minIndex] { // 每次遍历将i+1~len的数中选取最小的值 和当前的i节点交换
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func insertSort(arr []int) {
	len := len(arr)
	preIndex, current := 0, 0

	for i:=1; i<len; i++ {
		preIndex = i-1
		current = arr[i]
		// 上一个节点值 > 当前节点值 => 当前节点赋值为上一个节点的值，前移
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}

// 分为若干子序列，分别进行插入排序，为简单插入排序的改进版；优先选择较远的元素，又名缩小增量排序
// 时间复杂度O(NlogN)
func shellSort(arr []int) {
	len := len(arr)
	for gap:=len/2; gap>0; gap=gap/2 {
		for i:=gap; i<len; i++ {
			j:=i
			current := arr[i]
			for j-gap>=0 && current<arr[j-gap] {
				arr[j] = arr[j-gap]
				j = j-gap
			}
			arr[j] = current
		}
	}
}
