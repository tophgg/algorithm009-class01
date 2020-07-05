学习笔记

###十三课、字典树和并查集

####字典树Trie
主要应用于统计和排序大量的字符串（不限于字符串），所以经常被搜索引擎系统用于文本词频统计

优点：最大限度地减少无谓字符串的比较，查询效率比哈希表高

基本性质：

1.节点本身不存完整单词

2.从根结点到某一结点，路径上经过的字符连接起来，为该结点对应的字符串；

3.每个结点的所有子结点路径代表的字符都不相同（最终结点可以存放完结标识或者当前完整单词）


```
// Trie模板 go语言实现
type Trie struct {
	is_end bool
	next [26]*Trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	var node *Trie = this
	for _, ch := range word {
		if node.next[ch - 'a'] == nil {
			node.next[ch - 'a'] = new(Trie)
		}
		node = node.next[ch - 'a']
	}
	node.is_end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var node *Trie = this
	for _, ch := range word {
		node = node.next[ch - 'a']
		if node == nil {
			return false
		}
	}
	return node.is_end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var node *Trie = this
	for _, ch := range prefix {
			node = node.next[ch - 'a']
		if node == nil {
			return false
		}
	}
	return true
}
```

####并查集
主要应用于组团或配对问题，如朋友圈、岛屿数量等

```$xslt
// 并查集模板 go语言实现
// 连通分量，没连通2个，则减少一个联通分量
// 可用于朋友圈问题、底层变量的不同引用等
var Parent = make([]int, n)
var Count = 0
var Size = []int{}
func Init(n int) {
	Count = n
	Parent = make([]int, n)
	Size = make([]int, n)
	for i:=0; i<n; i++ {
		Parent[i] = i
		Size[i] = i
	}
}

// 返回某个节点p的根节点，主要复杂度，即待优化api
// 如果将这个树高度减为2，可以减少遍历次数
func Find(p int) int {
	// 根节点的parent[p] = p
	for p != Parent[p] {
		// 路径压缩，降为O(1)
		Parent[p] = Parent[Parent[p]]
		p = Parent[p]
	}
	return p
}

// 合并2个集合
// 可做平衡性优化，小的树接到大的树上，用size记录每个树的高度
func Union(p, q int) {
	rootP := Find(p)
	rootQ := Find(q)
	if rootP == rootQ {
		return
	}
	// 以其中一个集合为顶点
	//Parent[rootP] = rootQ
	if Size[rootP] > Size[rootQ] {
		Parent[rootQ] = rootP
		Size[rootP] += Size[rootQ]
	} else {
		Parent[rootP] = rootQ
		Size[rootQ] += Size[rootP]
	}

	Count--
}

func count() int {
	return Count
}
```



###十四课、高级搜索

剪枝的实现和特性：

双向BFS：从两侧同时遍历，适合BFS扩散比较大，这样可以减少深层BFS树的遍历

启发式搜索A*(基于BFS)：

曼哈顿距离： |x1-x2| + |y1-y2|

首先定义估价函数，可以根据问题来定义更少分治的路径函数，比如【1091二进制矩阵中的最短路径】，不断往斜下走，可以减少很多同层的路径；（对比BFS）

8Puzzles解法比较
RunningTime： 双向BFS A* < 双向BFS <=  A*(曼哈顿) < A*(汉明解法) << BFS


###十五课、红黑树和AVL树
常用搜索二叉树：
2-3 tree、
AVL tree 、红黑树、

由于二叉树的搜索时间主要取决于树的深度

####平衡二叉树：AVL Tree

通过定义平衡因子来维护平衡二叉树，减少树高，减少出现一条长链表的情况

平衡因子：左子树的高度减去右子树的高度 {-1, 0, -1}； 通过旋转操作来维护平衡（左旋、右旋、左右旋、右左旋）

不足：节点需要存储额外信息，且调整次数频繁（几乎每次增删都需要修改）

#### 近似平衡二叉树：红黑树
定义：近似平衡的二叉搜索树，保证任何一个结点的左右子树的高度差小于两倍。

特性：

1.每个结点要么是红色，要么是黑色

2.根节点是黑色

3.每个叶节点（NIL结点，空结点）是黑色的

4.不能有相邻的两个红色结点

5.从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点


```红黑树与AVL树比较

性能比较：

比较读性能，AVL比红黑更快，因为更平衡

改删，AVL更可能要维护节点，进行改动，而红黑因为有大于2倍的空间

AVL存储的信息更多，红黑树存的只需要一个bit来存是黑或红


实际应用：

读多写少用AVL

读和写一半一半  一般用红黑，逻辑更简洁，实现稍微简单些

读少写多用红黑树
```


课后作业分析：

【212.单词搜索II】用Trie实现的时间复杂度

1.传统Trie，从头开始遍历二维数组，每个字符串的4连通情况，确定一个字符后，只可选剩余的3个字符，不可回退
时间复杂度为0(M x 4 x 3^L-1)，空间复杂度为O(N)，其中M是单元格单词个数，N是待匹配单词数量，L是最长单词数量，如果优化了查找到一个单词就清除Trie中的节点，可减少重复判断。
