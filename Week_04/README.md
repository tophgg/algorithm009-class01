学习笔记

###第九课：深度优先搜素和广度优先搜索
Dfs，主要用栈来实现，从最底部开始

Bfs，主要用队列来实现，每层记录结果，直到最后

Dfs和bfs一般都会遍历所有的点，根据题目定义，可以进行一些剪枝处理，
或者双向bfs处理，或者dfs+bfs处理

```$xslt
# dfs模板
def dfs(node):
    if node in visited:
        # already visited
        return
    visited.add(node)
    # process current node
    # ... # logic here
    dfs(node.left)
    dfs(node.right)
```

```$xslt
# bfs模板
def bfs(node):
    queue = []
    queue.append([start])
    visited.add(node)
    while queue:
        node = queue.pop()
        visited.add(node)
        
        # process current node
        # ... # logic here
        process(node)
        nodes = generate_related_nodes
        queue.push(nodes)
        # other processing work
        ...
```

### 第十课：贪心算法

定义：贪心算法是在每一步选择中都采取在当前状态下最优或最好的选择，从而希望导致结果是全局最好或最优的算法
与动态规划不同在于它对每个子问题的解决方案都作出选择，不能回退。动态规划会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能。
贪心可以解决一些最优化问题。如图中的最小生成树。求哈夫曼编码等。然而对于工程和生活问题，贪心法一般不能得到我们锁要求的答案
一般较为高效和接近最优，可以用来辅助计算。

    例题：322硬币兑换

    如果是特殊的，比如硬币之间是整除关系比如【5、10、20】，可以直接用贪心算法得出结果，否则会匹配不出如【1、9、10】

    难点：证明可用贪心，贪心的套路。

###第十一课：二分查找

二分查找的使用条件：

1.具有单调性

2.具有上下界限

3.元素可以通过index进行查找替换

```$xslt
def bin-search(array, target):
    left, right = 0, len(array)-1
    while left <= right:
        mid = left + (right-left)/2
        if array[mid] == target:
            # find the target!
            break or return result
        elif array[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
```

算法拓展：牛顿迭代法
求切线无限逼近曲线
```
def sqrt(x):
    a = x
    for a * a > x:
        a = (a + x/a)/2
        
```
