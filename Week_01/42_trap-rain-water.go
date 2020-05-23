package main
// 1.暴力法 从1开始，向左循环遍历，每次都要从左到右找出左右侧的最小值来减去当前高度，累加每列的高度【弊端，每次都要遍历求左右最高的墙】
// 时间复杂度：O(n^2),空间复杂度：O(1)
// 2.动态规划 用两个数组来记录每列的左右侧的最高高度，遍历2次获得；最后遍历一次移动，可根据之前的栈来比较累加
// 时间复杂度：O(n),空间复杂度：O(n)
// 3.双指针 双向夹逼，每侧只相信对应侧的max值，计算max和当前列的差值，并更新当前侧的最高高度，并向中间聚拢
// 时间复杂度：O(n),空间复杂度：O(1)
// 4.栈 ...
func trap(height []int) int {
    left, right := 0, len(height)-1
    left_max, right_max, ans := 0, 0, 0

    for left <= right {
        // 对于左侧节点而言，left_max<right_max时为left_max为可信最小高度
        if left_max < right_max {
            ans += max(0, left_max - height[left])
            left_max = max(left_max, height[left])
            left++
        } else {
            ans += max(0, right_max - height[right])
            right_max = max(right_max, height[right])
            right--
        }
    }
    return ans
}

func max (a, b int) int {
    if a > b {
        return a
    }
    return b
}