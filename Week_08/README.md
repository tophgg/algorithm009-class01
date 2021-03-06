学习笔记

### 第十六课、位运算

基本操作：
```$xslt
左移 <<
    0011  =>  0110
右移 >>
    0110  =>  0011
或 |
    0011
    1011   =>  1011
与 & (11取1，10、00取0)
    0011
    1011   =>  0011
取反 ~
    0011   =>  1100
异或 ^ 相同为0不同为1
    0011
    1011    =>  1000
    
>>> 表示无符号右移，也叫逻辑右移，即若该数为正，则高位补0，而若该数为负数，则右移后高位同样补0
```

指定位置的位运算：
```$xslt
1.将x最右边的n位清零： 
    x & (~0 << n)   n左移n位得到10000再取与操作
2.获取x的第n位值（0或1）：
    (x >> n) & 1  右移n位后与1
3.获取x的第n位的幂值： 
    x & (1 << n)
4.仅将第n位置为1： 
    x | (1<<n)
5.仅将第n位置为0： 
    x & (~(1<<n))
6.将x最高位至第n位（含）清零： 
    x & ((1 << n) - 1)
```

实战使用：
```$xslt
一、判断奇偶：
x%2 == 1  —> (x&1) == 1
x%2 == 0 —> (x&1) == 0

二、除2： x/2  —> x >> 1
例mid=(left+right)/2  —> mid=(left+right)>>1

三、 清除最低位的1； x = x&(x-1)
  00100
& 00011

四、last = x & -x  => 得到最低位的1

五、清零 y = x & ~x  => 全0
```

### 第十七课、布隆过滤器和LRU缓存

布隆过滤器 对比 哈希表

哈希表 → hash后放入哈希表中，冲突后通过链表存储

布隆过滤器 → 只判断是否有 用二进制向量和随机映射函数存储，不存储具体数据，用于检索；

优点是空间效率和查询时间都远优于其他算法，缺点是删除困难，可能出现不精准；

布隆过滤器插入元素后（二进制位），在测试元素查询时，如果查到位都为1只能说可能存在，但是如果有一个位为0则肯定不存在该数据。

应用：做在缓存的最外级缓存，查询数据库前先查布隆过滤器，查看是否存在对应的数据（如果不存在则直接忽略该查询）；正常查询数据库后，写入布隆过滤器标记可以访问

实现：hash_num定位随机hash的次数，位数；初始化时均为0；根据多次hash，分别hash后写入二进制位，判断是否均为1或存在0，返回对应的结果

LRU缓存

科普：cpu三级缓存，速度：L1 > L2 > L3 ；存储大小：L1 < L2 < L3

lru cache的二要素：大小，替换策略（LRU、LFU）；  实现：Hash Table + Double LinkedList； O(1)的查询、修改、更新


###第十八课、排序算法

比较类排序 （重点看nlogn的排序算法）

十大排序算法：[https://www.notion.so/6e13f47548ad49bda6509a127b2151ce#dea65311d02a4c8a930cd4d728eb2b08](https://www.notion.so/6e13f47548ad49bda6509a127b2151ce#dea65311d02a4c8a930cd4d728eb2b08)

- 交换排序 —
    - 冒泡：嵌套循环，每次查看相连元素，如果逆序则交换 O(N^2)
    - *快排：选定一个标杆pivot（可优化，随机靠中间），小元素放左边，大元素放右边，不断递归调用左右元素，继续分治。模板，无额外空间调用 - 练习
- 插入排序  —
    - 简单插入排序：从前到后逐步构建有序序列，对未排序数据，在已排序序列中从后往前扫描，找到相应的位置插入 O(N^2)
    - 希尔排序
- 选择排序  —
    - 简单选择排序：每次选择最小值，然后放到待排序数组的起始位置 O(N^2)
    - *堆排序：堆插入O(logN)，取最大小值O(1)；数组元素一次建立小顶堆；依次取堆顶元素，并删除。二叉堆：完整二叉树可用数组表示（从0开始左儿子：2i+1，右儿子2i+2）
- *归并排序  —
    - 二路归并排序：将数组分为2两块，分别求左右的排序，最后将2路有序数组进行合并
    - 多路归并排序

非比较类排序

- 计数排序
- 桶排序
- 基数排序

各种情况下排序算法的选择：

完全随机：

快排 = 归并 ≥ 堆排 > 希尔排序 > 插入 = 选择 = 冒泡

有重复元素：

希尔排序 = 归并 = 堆 ≥ 快排 > 插入 > 选择 = 冒泡

逆序元素：

归并 = 堆排 = 希尔 > 插入 = 选择 = 冒泡 = 快排（退化为n^2）

基本有序元素：

插入 = 冒泡 = 希尔 = 归并 > 堆排 > 选择 = 快排