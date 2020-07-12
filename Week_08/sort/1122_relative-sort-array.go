package main

import "sort"

// 桶排序，事先生成可能存放数据大小的桶， 牺牲了空间，但是速度会快
// 时间复杂度O(KN)，空间复杂度O(N)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	bucket := make([]int, 1001)
	// 先记录arr1的元素，并计数
	for _, v := range arr1 {
		bucket[v]++
	}
	// 将arr2的元素，先排在arr1的前面，按之前计数的方式把arr1相同的元素都排放好
	i := 0
	for _, v := range arr2 {
		for bucket[v] > 0 {
			arr1[i] = v
			i++
			bucket[v]--
		}
	}
	// 将arr1剩余的元素再继续排列 -- 因为桶里的数值都是已经排好序的了 按升序存放！！
	for v, _ := range bucket {
		for bucket[v] > 0 {
			arr1[i] = v
			i++
			bucket[v]--
		}
	}
	return arr1
}


// 计数排序，时间复杂度：O(KN + NlogN), 空间复杂度：O(N)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	hash := make(map[int]int, 0)
	res := []int{}
	for _, v := range arr1 {
		hash[v]++
	}

	i := 0
	for _, v := range arr2 {
		for hash[v] > 0 {
			res = append(res, v)
			hash[v]--
			i++
		}
	}

	tmp := make([]int, 0)
	for k, v := range hash {
		for v > 0 {
			tmp = append(tmp, k)
			v--
		}
	}
	sort.Ints(tmp)
	return append(res, tmp...)
}