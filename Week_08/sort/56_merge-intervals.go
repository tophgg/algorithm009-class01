package main

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 排序前有6种情况，排序后只有4种
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}

	begin := intervals[0][0]
	end := intervals[0][1]
	for i:=1; i<len(intervals); i++ {
		// 原右节点比新的左节点小，可以无交集，结果加入
		if end < intervals[i][0] {
			res = append(res, []int{begin, end})
			begin = intervals[i][0]
			end = intervals[i][1]
		} else {
			// 取原右节点和新右节点的最大值
			end = max(end, intervals[i][1])
		}
	}
	res = append(res, []int{begin, end})
	return res
}

func max(a, b int) int {
	if a > b {return a}
	return b
}
