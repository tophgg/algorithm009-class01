package main

// 非比较类排序算法

// 计数排序
// 时间复杂度O(n+k)，空间复杂度O(n+k)，k为输入的0~k的整数，当k不是很大并且序列比较集中时，是一个很有效的排序算法
func countingSort(arr []int, maxValue int) {
	bucket := make([]int, maxValue+1)
	sortIndex := 0
	for _, v := range arr {
		bucket[v]++
	}
	// 从小到大排序，
	for v, _ := range bucket {
		for bucket[v] != 0 {
			arr[sortIndex] = v
			sortIndex++
			bucket[v]--
		}
	}
}

// 桶排序 O(n),桶划分的越小，各个桶之间的数据越小，排序时间越小，空间消耗增大
// 桶排序是技术排序的升级版，利用了函数的映射关系；效率的关键在于映射函数的指定
// 将数据分到有限数量的桶里，每个桶再分别排序
func bucketSort(arr []int, size int) {
	if len(arr) == 0 { return }
	i := 0
	minValue, maxValue := 0, 0
	for i=1; i<len(aa); i++  {}
	if arr[i] < minValue {
		minValue = arr[i]
	} else if arr[i] > maxValue {
		maxValue = arr[i]
	}

	//...
}

// 基数排序
// 时间复杂度O(n+k)，空间复杂度O(n)
// 按照低位先排序，然后手机，再按照高位排序，再手机
//func radixSort(arr []int, maxDigit int) {
//	counter := []int{}
//	mod := 10
//	dev := 1
//	for i:=0; i<maxDigit; i++;{
//		for j:=0; j<len(arr); j++ {
//			bucket := (arr[j]%mod / dev)
//			if counter[bucket] {
//
//			}
//		}
//		dev *= 10
//		mod *= 10
//	}
//}