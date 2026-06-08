package main

import "time"

/*
实现一个支持过期时间（TTL）的 LRU 缓存。

缓存中的每个 key 在写入时都会指定一个 TTL（单位：秒）。

当 key 超过 TTL 后，视为失效。

缓存容量固定，采用 LRU 淘汰策略。

要求实现以下接口：
*/

type LRUCache struct {
	cache          map[int]*DLinkedNode // 哈希表
	size, capacity int                  // 当前元素数量和总容量
	head, tail     *DLinkedNode         // 头尾指针
}

type DLinkedNode struct { // 双向链表
	key, value int
	expireTime time.Time
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*DLinkedNode),
		size:     0,
		capacity: capacity,
		head:     new(DLinkedNode),
		tail:     new(DLinkedNode),
	}

	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (this *LRUCache) Put(key int, value int, ttl int) {
	// 先看 key 存不存在
	if _, exists := this.cache[key]; exists { // key 存在
		node := this.cache[key]
		now := time.Now()
		if now.Before(node.expireTime) { // 没有过期
			// 更改值和过期时间
			node.value = value
			node.expireTime = now.Add(time.Duration(ttl) * time.Second)

			// 将节点移到最前面
			this.moveToHead(node)
			return
		} else { // 过期了，删除旧节点，走新建节点的流程
			this.removeNode(node)
			this.size--
			delete(this.cache, key)
		}
	}

	// 新的 key
	node := &DLinkedNode{
		key:        key,
		value:      value,
		expireTime: time.Now().Add(time.Duration(ttl) * time.Second),
	}
	this.cache[key] = node

	// 插入成为首节点
	this.addToHead(node)
	this.size++

	// 可能导致超出容量
	if this.size > this.capacity {
		// 删除尾部的节点
		move := this.tail.prev
		this.removeNode(move)
		this.size--
		delete(this.cache, move.key)
	}
}

func (this *LRUCache) Get(key int) int {
	now := time.Now()
	node, exists := this.cache[key]

	if !exists { // 不存在，直接返回-1
		return -1
	}

	if !now.Before(node.expireTime) { // 存在但过期，删除节点
		this.removeNode(node)
		this.size--
		delete(this.cache, key)

		return -1
	}

	// 存在且未过期，将节点放到前面，返回 value
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
