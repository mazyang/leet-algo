package main

// LeetCode 146
// Least Recently Used(最近最少使用)
// LRU算法的思想是：如果一个数据在最近一段时间没有被访问到，那么可以认为在将来它被访问的可能性也很小。
// 因此，当空间满时，最久没有访问的数据最先被置换（淘汰）
// LRU 把数据按照时间排序

// Node 双链表节点定义
type Node struct {
	Key, Value int
	Prev, Next *Node
}

// DoubleList 因为需要删除节点，所以使用双向链表保证时间复杂度为 O(1)
// 靠近队尾的节点是最近使用的，越靠近头部的节点就是越久未使用的
type DoubleList struct {
	// 头尾虚节点
	Head, Tail *Node
	// 链表节点数
	Size int
}

// DoubleListConstruct 双向链表构造器
func DoubleListConstruct() *DoubleList {
	dl := &DoubleList{}
	dl.Head = &Node{Key: 0, Value: 0}
	dl.Tail = &Node{Key: 0, Value: 0}
	dl.Head.Next = dl.Tail // 并不是循环链表
	dl.Tail.Prev = dl.Head
	dl.Size = 0
	return dl
}

// 在链表尾部添加节点
func (dl *DoubleList) addLast(x *Node) {
	x.Prev = dl.Tail.Prev
	x.Next = dl.Tail
	dl.Tail.Prev.Next = x
	dl.Tail.Prev = x
	dl.Size++
}

// 删除指定节点
func (dl *DoubleList) remove(x *Node) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
	dl.Size--
}

// 删除链表的第一个节点
func (dl *DoubleList) removeFirst() *Node {
	if dl.Head.Next == dl.Tail {
		return nil
	}
	first := dl.Head.Next
	dl.remove(first)
	return first
}

// LRUCache 使用哈希链表（双向链表+哈希表）作为底层数据结构
type LRUCache struct {
	// key -> Node(key, val)
	hashMap map[int]*Node // 通过key快速查找到一个Node
	// Node(key1, val1) <-> Node(key2, val2)
	cache *DoubleList // 双向链表支持快速插入和删除第一个和最后一个节点
	// 最大容量
	cap int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{cap: capacity}
	lru.hashMap = make(map[int]*Node)
	lru.cache = DoubleListConstruct()
	return lru
}

// 在两个数据结构之上提供抽象的API，避免Get和Put直接操作 hashMap 和 cache 的细节

// 将某个key提升为最近使用的（靠近尾部的节点是最近使用的）
func (this *LRUCache) makeRecently(key int) {
	x := this.hashMap[key]
	// 先从链表中删除key对应的节点
	this.cache.remove(x)
	// 重新插入到队尾
	this.cache.addLast(x)
}

// 添加最近使用的元素
func (this *LRUCache) addRecently(key, val int) {
	x := &Node{Key: key, Value: val}
	this.cache.addLast(x)
	this.hashMap[key] = x
}

// 根据key删除一个元素
func (this *LRUCache) deleteNode(key int) {
	x := this.hashMap[key]
	this.cache.remove(x)
	delete(this.hashMap, key)
}

// 删除最近未使用的元素
func (this *LRUCache) removeLeastRecently() {
	// 链表头部第一个元素就是最久未使用的
	x := this.cache.removeFirst()
	// 从hashMap中删除
	delete(this.hashMap, x.Key)
}

func (this *LRUCache) Get(key int) int {
	// 1. 如果不存在，返回-1
	if _, ok := this.hashMap[key]; !ok {
		return -1
	}
	// 2. 如果存在，设置为最近访问的
	this.makeRecently(key)
	return this.hashMap[key].Value
}

func (this *LRUCache) Put(key int, value int) {
	// 如果已经存在，更新value
	if _, ok := this.hashMap[key]; ok {
		this.deleteNode(key)
		this.addRecently(key, value)
		return
	}
	// 如果容量已满, 先删除最近未使用的，再添加
	if this.cap == this.cache.Size {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}
