package offer

// 双向链表

type Snode struct {
	prev *Snode
	next *Snode
	key  int
	val  int
}
type LRUCache struct {
	// key: value
	mp   map[int]*Snode
	cap  int
	head *Snode
	tail *Snode
}

func Constructor3(capacity int) LRUCache {
	head := &Snode{val: -1}
	tail := &Snode{val: -1}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:  capacity,
		mp:   map[int]*Snode{},
		head: head,
		tail: tail,
	}
}

func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.mp[key]; ok {
		lru.Put(key, lru.mp[key].val)
		return v.val
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key int, value int) {
	v, ok := lru.mp[key]
	if ok {
		v.val = value
		// 删除原先关系
		v.prev.next = v.next
		v.next.prev = v.prev
		// 连接到head头部
		v.prev = lru.head
		v.next = lru.head.next
		lru.head.next.prev = v
		lru.head.next = v
	} else {
		// 如果不存在
		if len(lru.mp) == lru.cap {
			// 获取last
			tail := lru.tail.prev
			// 删除last
			tail.prev.next = tail.next
			tail.next.prev = tail.prev
			delete(lru.mp, tail.key)
		}
		// insert
		newNode := &Snode{key: key, val: value}
		lru.mp[key] = newNode
		// 连接到head头部
		newNode.prev = lru.head
		newNode.next = lru.head.next
		lru.head.next.prev = newNode
		lru.head.next = newNode
	}
}
