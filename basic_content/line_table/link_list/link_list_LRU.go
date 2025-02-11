package link_list

import (
	"fmt"
	"strings"
)

// 使用双向链表实现LRU淘汰策略
// get平均时间复杂度：O(n)
// put平均时间复杂度：O(n)

// 缓存常用淘汰策略
// 1. 先进先出策略 FIFO（First In，First Out）
// 2. 最少使用策略 LFU（Least Frequently Used）
// 3. 最近最少使用策略 LRU（Least Recently Used）

type LRUCache struct {
	capacity int
	len      int
	// 带头链表
	head *CacheNode
	tail *CacheNode
}

func (c *LRUCache) ToString() string {
	if c.isEmpty() {
		return ""
	}

	results := make([]string, 0, c.len)
	node := c.head.next
	for node != c.tail {
		results = append(results, node.ToString())
		node = node.next
	}

	return strings.Join(results, "->")
}

func (c *LRUCache) moveToHead(node *CacheNode) {
	if c.isEmpty() {
		return
	}

	// 是否是头
	if node == c.head.next {
		return
	}

	// 是否是尾部
	if node == c.tail.prev {
		node.prev.next = c.tail
		node.next = c.head.next
		c.tail.prev = node.prev
		node.prev = c.head
		c.head.next.prev = node
		c.head.next = node
		return
	}

	// 取出node
	node.next.prev = node.prev
	node.prev.next = node.next

	// 置于头部
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) getNode(key string) *CacheNode {
	if c.isEmpty() {
		return nil
	}

	i := c.head.next
	for i.next != c.tail {
		if i.key == key {
			return i
		}

		i = i.next
	}

	return nil
}

func (c *LRUCache) contains(key string) bool {
	return c.getNode(key) != nil
}

func (c *LRUCache) PutValue(key string, value int) {
	if c.contains(key) {
		return
	}

	if c.isFull() {
		// 淘汰尾节点
		c.removeLast()
	}

	node := &CacheNode{
		key:   key,
		value: value,
	}

	if c.isEmpty() {
		c.head.next = node
		c.tail.prev = node
		node.next = c.tail
		node.prev = c.head
	} else {
		c.head.next.prev = node
		node.next = c.head.next
		node.prev = c.head
		c.head.next = node
	}
	c.len++
}

func (c *LRUCache) removeLast() {
	if c.isEmpty() {
		return
	}

	if c.len == 1 {
		c.head.next = c.tail
		c.tail.prev = c.head
	} else {
		c.tail.prev = c.tail.prev.prev
		c.tail.prev.next = c.tail
	}

	c.len--
}

func (c *LRUCache) GetValue(key string) *CacheNode {
	node := c.getNode(key)
	if node == nil {
		return nil
	}

	c.moveToHead(node)
	return node
}

func (c *LRUCache) isFull() bool {
	return c.len == c.capacity
}

func (c *LRUCache) isEmpty() bool {
	return c.len == 0
}

func Constructor(capacity int) *LRUCache {
	if capacity <= 0 {
		return nil
	}

	// 创建虚拟头节点, 哨兵节点
	head := &CacheNode{}
	tail := &CacheNode{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

type CacheNode struct {
	key   string
	value int
	prev  *CacheNode
	next  *CacheNode
}

func (n *CacheNode) ToString() string {
	return fmt.Sprintf("key:%s,value:%d", n.key, n.value)
}

func (n *CacheNode) getValue() int {
	return n.value
}
