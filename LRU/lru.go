package main

//// LRUCache 使用哈希链表（双向链表+哈希表）作为底层数据结构
//type LRUCache struct {
//	// key -> Node(key, val)
//	hashMap map[int]*list.Element
//	// Node(key1, val1) <-> Node(key2, val2)
//	cache *list.List
//	// 最大容量
//	cap int
//}
//
//func Constructor(capacity int) LRUCache {
//	lru := LRUCache{cap: capacity}
//	lru.hashMap = make(map[int]*list.Element)
//	lru.cache = list.New()
//	return lru
//}
//
//// 在两个数据结构之上提供抽象的API，避免Get和Put直接操作 hashMap 和 cache 的细节
//
//// 将某个key提升为最近使用的
//func (this *LRUCache) makeRecently(key int) {
//	x := this.hashMap[key]
//	// 先从链表中删除key对应的节点
//	this.cache.Remove(x)
//	// 重新插入到队尾
//	this.cache.PushBack(x.Value.(*Node))
//}
//
//// 添加最近使用的元素
//func (this *LRUCache) addRecently(key, val int) {
//	x := &Node{Key: key, Value: val}
//	this.cache.PushBack(x)
//	this.hashMap[key] = this.cache.Back()
//}
//
//// 根据key删除一个元素
//func (this *LRUCache) deleteKey(key int) {
//	x := this.hashMap[key]
//	this.cache.Remove(x)
//	delete(this.hashMap, key)
//}
//
//// 删除最近未使用的元素
//func (this *LRUCache) removeLeaseRecently() {
//	// 链表头部第一个元素就是最久未使用的
//	e := this.cache.Front()
//	x := this.cache.Remove(e)
//	// 从hashMap中删除
//	delete(this.hashMap, x.(*Node).Key)
//}
//
//func (this *LRUCache) Get(key int) int {
//	if _, ok := this.hashMap[key]; !ok {
//		return -1
//	}
//	this.makeRecently(key)
//	return this.hashMap[key].Value.(*Node).Value
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	// 如果已经存在，更新value
//	if _, ok := this.hashMap[key]; ok {
//		this.deleteKey(key)
//		this.addRecently(key, value)
//		return
//	}
//	// 如果容量已满
//	if this.cap == this.cache.Len() {
//		this.removeLeaseRecently()
//	}
//	this.addRecently(key, value)
//}
//
//func main() {
//	lRUCache := Constructor(2)
//	lRUCache.Put(1, 1)
//	lRUCache.Put(2, 2)
//	fmt.Println(lRUCache.Get(1)) // 返回 1
//	lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
//	lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//	fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
//	fmt.Println(lRUCache.Get(3)) // 返回 3
//	fmt.Println(lRUCache.Get(4)) // 返回 4
//}
