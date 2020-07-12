package main

func mergeSort(arr []int, left, right int) {
	if right <= left {return}
	mid := (left+right) >> 1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i<=mid && j<= right {
		if arr[i] > arr[j] {
			tmp[k] = arr[j]
			k++;j++
		} else {
			tmp[k] = arr[i]
			k++; i++
		}
	}
	for i<=mid {
		tmp[k] = arr[i]
		k++; i++
	}
	for j<=right {
		tmp[k] = arr[j]
		k++; j++
	}
	for p:=0; p<len(tmp); p++ {
		arr[left+p] = tmp[p]
	}
}