package code

// 定义一个node
type LFUNode struct {
	key   int
	value int
	hits  int
	prev  *LFUNode
	next  *LFUNode
}

type LFUCache struct {
	capacity int
	used     int
	cache    map[int]*LFUNode
	head     *LFUNode
	tail     *LFUNode
}

func NewLFUCache(capacity int) LFUCache {
	lru := LFUCache{
		capacity: capacity,
		used:     0,
		cache:    make(map[int]*LFUNode),
		head:     &LFUNode{},
		tail:     &LFUNode{},
	}
	//建立双向链表
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LFUCache) Get(key int) int {

	node, ok := lru.cache[key]
	if ok {
		//移动最表头
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

func (lru *LFUCache) Put(key, value int) {
	//如果已存在,直接更新，并移动到表头
	node, ok := lru.cache[key]
	if ok {
		node.value = value
		lru.moveToHead(node)
		return
	}

	newNode := &LFUNode{
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
		//移除链毛
		removed := lru.removeTail()
		//删除cache
		delete(lru.cache, removed.key)
		lru.used--
	}
}

func (lru *LFUCache) addToHead(node *LFUNode) {
	//当前节点的prev指向head
	node.prev = lru.head
	//当前节点的next指向head的next
	node.next = lru.head.next
	//head节点的next的prev指向当前node
	lru.head.next.prev = node
	//head节点的next指向node
	lru.head.next = node
}

func (lru *LFUCache) removeNode(node *LFUNode) {
	//当前节点的prev的next指向当前节点的next
	node.prev.next = node.next
	//当前节点的next的prev指向当前节点的prev
	node.next.prev = node.prev
}

func (lru *LFUCache) moveToHead(node *LFUNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LFUCache) removeTail() *LFUNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
