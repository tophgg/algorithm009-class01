学习笔记

##第十九课、高级动态规划
复习模板
递归、分治、回溯


##第二十课、字符串算法

不可变字符串 - 线程安全

如go、java、python

字符串比较

go：...

字符串遍历

go：...

滑动窗口模板：
```
s := "xxxxx"
p := "xx"
l, r := 0, 0
valid := 0
window := make(map[byte]int)
need := make(map[byte]int)
res := make([]int)
for r<len(s) {
    c := s[r]
    r++
    if _, ok := need[c]; ok {
        window[c]++
        if window[c] == need[c] {
            valid++
        }
    }
    // some logic to move window
    for r-l <= len(p) {
        if len(need) == valid {
            // collect
            res = append(res, l)
        }
        d := s[l]
        l++
        if _, ok := need[d]; ok {
            if window[d] == need[d] {
                valid--
            }
            window[d]--
        }
    }
    return res
}

```

动态规划+字符串

字符串匹配算法

子串长度为M，目标串长度为N

1.暴力法 O(MN)，不断移动每次1个位置逐个对比，两层嵌套循环

2.Rabin-Karp算法

优化hash，使用hash对比2个子串是否相等，不用逐个移动，滑动窗口

256进制位运算，保证hash只需要O(1)的复杂度来比较，不然hash(M)如果有M的复杂度

最坏O(MN)，一般O(N)复杂度

3.KMP算法

优化暴力，求出最长前缀和和后缀和，当匹配失败后，每次移动n个位置找到下个匹配的前缀

讲解视频：[bilibili.com/video/av11866460?from=search&seid=17425875345653862171](http://bilibili.com/video/av11866460?from=search&seid=17425875345653862171)

廖雪峰讲解：[http://www.ruanyifeng.com/blog/2013/05/Knuth–Morris–Pratt_algorithm.html](http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html)

了解-Sunday算法