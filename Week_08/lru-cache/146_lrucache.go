package list


type LRUCache struct {
	cap    int   // 缓存总数
	header *Node // 头链表，用于记录最近使用的节点
	tail   *Node // 尾链表，用于记录最后一个节点
	m      map[int]*Node // 借助map存放链表节点，来实现查找
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap: capacity,
		header: &Node{},
		tail: &Node{},
		m: make(map[int]*Node, capacity),
	}
	// 初始化头尾链表，让他们互相指向对方
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}


func (this *LRUCache) Get(key int) int {
	// 查看节点是否在map中
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		// 存在的话需要把这个node移动到header放到最上面，不关心是否为第一个
		// 1.把当前节点从链表中去除 2.在头链表指向当前节点 3.返回当前节点的value
		this.remove(node)
		this.putHead(node)
		return node.value
	}
}

//
func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.m[key]; ok {
		//存在节点，更新数据，移动到head
		node.value = value
		this.remove(node)
		this.putHead(node)
	} else {
		// 不存在
		// 如果容器已经满了，需要删除链表尾部，从map中删除
		// part1:
		//  m[delKey]
		// ... pre <-> tail
		if len(this.m) >= this.cap {
			deleteKey := this.tail.pre.key
			this.remove(this.tail.pre)
			delete(this.m, deleteKey)
		}

		// 创建新的节点，并放在head，添加到map中
		newNode := Node{
			key: key,
			value: value,
		}
		this.putHead(&newNode)
		this.m[key] = &newNode
	}
}

// 删除当前的node节点，让上一个节点的next指针指向当前节点的next指针，
// 让下一个节点的pre指针指向当前节点的pre指针
func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//  将当前节点放到头节点去
func (this *LRUCache) putHead(node *Node) {
	// part1:
	//            ->          node              ->          pre     next
	//  node             head <-> originNext          head <-> node <-> originnEXT ...
	originNext := this.header.next
	this.header.next = node
	node.next = originNext

	originNext.pre = node
	node.pre = this.header
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
