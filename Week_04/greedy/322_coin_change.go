package main


import (
	"math"
	"sort"
)

// 解法1：贪心+dfs
// 解法2：dp
func coinChange(coins []int, amount int) int {
	minCount := math.MaxInt32
	sort.Ints(coins)
	// 从后往前，贪心
	dfs(coins,len(coins)-1,amount,0,&minCount)
	if minCount == math.MaxInt32{
		return -1
	}
	return minCount
}

func dfs(coins []int, index int, amount int, count int, minCount *int){
	// index游标递减
	if index < 0{
		return
	}
	// 从大硬币到小硬币dfs
	coin := coins[index]
	// cnt为当前硬币的余数，也是当前硬币使用个数
	for cnt:=amount/coin; cnt>=0; cnt--{
		remainAmount := amount - coin * cnt
		curCount := count + cnt
		// 求尽得0，比较当前计数和历史最小计数比较
		if remainAmount == 0{
			*minCount = Min(*minCount,curCount )
			break
		}
		// 剪枝，比最小还多次数则中断
		if curCount +1 >= *minCount  {
			break
		}
		dfs(coins,index-1,remainAmount,curCount,minCount)
	}
}

func Min(a,b int) int {
	if a < b{
		return a
	}
	return b
}