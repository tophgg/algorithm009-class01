学习笔记

## 第五课：哈希表、映射、集合

哈希表-散列函数，可以把值（index）映射对应的值（interface{}）
如果哈希函数不够优，会出现碰撞冲突，有重复元素，哈希函数有几种拓展方法，如值变成链表，拉链式解决冲突法（有可能变成O(N)，链表），或者再hash一次。

Map-映射，key不重复

Set：不重复的元素集合

## 第六课：树、二叉树、二叉搜索树

链表是特殊化的树，只有一边的节点
树和图的区别，有无环

决策树、状态树、博弈树，决定复杂度

树的遍历：
前序遍历：根打印-左-右

中序遍历： 左-根打印-右

后序遍历：左-右-根打印

    二叉树：O(logN)，最坏O(N)退化为链表
    1.左子树所有节点的小于他的根节点
    2.右子树的所有节点都大于他的根节点
    3.左右子树都是如此，具有重复性
    操作：
    查询、删除、插入


####AVL树
是带有平衡条件的二叉查找树，所有节点的左右树高不超过1，无论是删除还是插入操作，只要不满足这个条件就需要通过旋转来重新保持平衡，比较耗时，适用于插入、查找少，查询多的情况。
如windows NI内核中广泛存在

####红黑树
弱平衡二叉树，旋转次数少，插入最多两次旋转，删除最多三次旋转。通过在每个节点增加一个存储位标识节点的颜色，可以是红或黑。通过对任意挑根到叶子的路径上各个节点着色的方式的限制，红黑树确保没有一条路径会比其他路径长出两倍。
如c++的STL中，map和set底层用红黑树实现。


堆：
堆的实现：有多种，如
二叉堆Binary、
Leftist左倾堆、
Binomial	、
F斐波那契ibonacci、
Pairing、
Brodal、
Rank-pairing、
Strict Fibonacci、
2-3heap


####二叉堆
    二叉堆的特性：
    1.一个完全二叉树
    2.任意节点都大于其根节点
    3.用一个数组实现

    二叉堆的查找特性：
    0.根节点顶堆是a[0]
    1.索引为i的左孩子的索引是(2*i+1)
    2.索引为i的右孩子的索引是(2*i+2)
    3.索引为i的父节点的索引是floor((i-1)/2); （取整）

    二叉堆的插入特性：
    1.新元素一律先插到堆的尾部
    2.依次向上调整整个堆的结构
    树的深度就是O(logn)就是调整的次数，而且每次上移时肯定会大于他的子节点。不需要其他操作

    二叉堆的删除特性：
    1.将堆尾元素替换到顶部（即对顶被替代删除掉）
    2.依次从根部向下调整整个堆的结构（一直到堆尾即可）*每次下移将左右节点较大的节点替换，直到没有子节点

拓展：查看go中"container/heap”的实现