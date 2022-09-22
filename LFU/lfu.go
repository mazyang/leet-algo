package main

// Least Frequently Used：最近不经常使用
// 根据访问缓存的历史频率来淘汰数据，核心思想是“如果数据在过去一段时间被访问的次数很少，那么将来被访问的概率也会很低”
// LFU 把数据按照访问频次排序，淘汰访问频次最低的数据，如果访问频次最低的数据有多条，需要淘汰最旧的数据(最早插入的数据)

type LFUCache struct {
	cache    map[int]*node // 通过key快速找到一个节点
	freqMap  map[int]*list // key 为访问频率，value为节点链表，可以通过 minFreq 找到最少访问频率的所有节点
	minFreq  int           // 最少访问的频率
	capacity int           // 最大容量
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*node),
		freqMap:  make(map[int]*list),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.incrFreq(n)
	return n.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// 如果已经存在，更新访问频率
	n, ok := this.cache[key]
	if ok {
		n.val = value
		this.incrFreq(n)
		return
	}
	// 如果达到最大容量
	if len(this.cache) == this.capacity {
		l := this.freqMap[this.minFreq]
		delete(this.cache, l.removeBack().key)
	}
	n = &node{key: key, val: value, freq: 1}
	this.addNode(n)
	this.cache[key] = n
	this.minFreq = 1
}

// 增加指定节点的访问频率
func (this *LFUCache) incrFreq(n *node) {
	l := this.freqMap[n.freq]
	l.remove(n)
	if l.empty() {
		delete(this.freqMap, n.freq)
		if n.freq == this.minFreq {
			this.minFreq++ // 直接增加minFreq即可，不需要遍历获得minFreq
		}
	}
	n.freq++
	this.addNode(n)
}

func (this *LFUCache) addNode(n *node) {
	l, ok := this.freqMap[n.freq]
	// 如果没有指定的访问频率的链表，就创建一个链表存放对应频率的所有节点
	if !ok {
		l = newList()
		this.freqMap[n.freq] = l
	}
	l.pushFront(n)
}

// 双向链表中的节点
type node struct {
	key  int
	val  int
	freq int // 访问频率
	prev *node
	next *node
}

// 双向链表
type list struct {
	head *node
	tail *node
}

func newList() *list {
	head := new(node)
	tail := new(node)
	head.next = tail
	tail.prev = head
	return &list{
		head: head,
		tail: tail,
	}
}

func (l *list) pushFront(n *node) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
}

func (l *list) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
}

func (l *list) removeBack() *node {
	n := l.tail.prev
	l.remove(n)
	return n
}

func (l *list) empty() bool {
	return l.head.next == l.tail
}
