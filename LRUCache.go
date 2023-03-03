package code

// 定义一个node
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	used     int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		used:     0,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	//建立双向链表
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {

	node, ok := lru.cache[key]
	if ok {
		//移动到表头
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	//如果已存在,直接更新，并移动到表头
	node, ok := lru.cache[key]
	if ok {
		node.value = value
		lru.moveToHead(node)
		return
	}

	newNode := &Node{
		key:   key,
		value: value,
	}
	//添加到缓存
	lru.cache[key] = newNode
	//添加到表头
	lru.addToHead(newNode)
	lru.used++
	//判断是否超时容量
	if lru.used > lru.capacity {
		//移除链尾
		removed := lru.removeTail()
		//删除cache
		delete(lru.cache, removed.key)
		lru.used--
	}
}

func (lru *LRUCache) addToHead(node *Node) {
	//当前节点的prev指向head
	node.prev = lru.head
	//当前节点的next指向head的next
	node.next = lru.head.next
	//head节点的next的prev指向当前node
	lru.head.next.prev = node
	//head节点的next指向node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *Node) {
	//当前节点的prev的next指向当前节点的next
	node.prev.next = node.next
	//当前节点的next的prev指向当前节点的prev
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *Node {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
