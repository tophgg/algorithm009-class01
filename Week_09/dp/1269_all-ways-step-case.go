package main

// 边界【有限步长+不能在0的左边】，cur==0&&steps==0 then return 1 && 0 < cur_index < arrLen
// 解法1：递归，确定重复的每一步，到达边界时退出递归循环，如果不优化，时间复杂度为O(3^N)，空间复杂度O(3^N)
// 解法2：递归+记忆存储，每次分支出去，其实会走重复的路线，优化存储每一个(step,cur_index)，可以保证，当前位置重复走过的路线都缓存，这样时间复杂度为O(N*M)，空间复杂度O(N*M)
// 解法3：动态规划，暴力版；这种二维数组，首先会有很多前面的重复项来得出 O(N^2)
//  1.定义状态f[i][j]表示第i步，指针位置j的方案数；
//  2.转移方程即当前这一步可从上一步得出：f[i][j] = f[i-1][j-1] + f[i-1][j] + f[i-1][j]
//  3.初始状态和边界，就是f[0][0] = 只有1种，j-1不能小于0，j+1不能超过steps

// 解法4：动态规划-优化-剪枝；这个二维数组，其实不需要重复那么多次，因为最多只能走steps步(<500)，所以一级数组的初始长度为min(arrLen, steps)
// 时间复杂度：O(steps*min(arrLen,steps+1))；空间复杂度：O(steps*min(arrLen,steps+1))
// 解法5：动态规划-滚动数组：维护两个个数组，用新旧数组来记录每次移动时上一次的数据即可。
// 时间复杂度：O(N)；空间复杂度：O(N)
// 解法6：动态规划-单一滚动数组：维护一个数组+上次的步数即可
// 时间复杂度：O(N)；空间复杂度：O(N)

// 解法1：递归暴力法
//func numWaysByLoop(steps int, arrLen int) int {
//	if steps <= 0 || arrLen <= 0 {
//		return 0
//	}
//	var res = dfsByLoop(steps, 0, arrLen)
//	return res % (1000000000+7)
//}
//
//func dfsByLoop(steps int, cur int, arrLen int) int {
//	if steps == 0 {
//		// 回到原点
//		if cur == 0 {
//			return 1
//		} else {
//			return 0
//		}
//	}
//
//	if cur < 0 || cur >= arrLen {
//		return 0
//	}
//
//	return dfsByLoop(steps - 1, cur + 1, arrLen) +
//		dfsByLoop(steps - 1, cur, arrLen) +
//		dfsByLoop(steps - 1, cur - 1, arrLen)
//}

// 解法2：递归+记忆存储

//class Solution:
//	def numWays(self, steps: int, arrLen: int) -> int:
//		memo = {}
//		return self.dfs(steps, 0, arrLen, memo)
//
//	def dfs(self, steps, cur, arrLen, memo):
//		if steps == 0:
//			return 1 if cur == 0 else 0
//		if cur < 0 or cur >= arrLen:
//			return 0
//		if (steps, cur) in memo:
//			return memo[(steps, cur)]
//		memo[(steps, cur)] = self.dfs(steps-1, cur-1, arrLen, memo)+self.dfs(steps-1, cur, arrLen, memo)+self.dfs(steps-1, cur+1, arrLen, memo)
//		return memo[(steps, cur)] % (10 ** 9 + 7)
