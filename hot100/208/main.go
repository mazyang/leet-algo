package main

type Trie struct {
	children [26]*Trie
	isEnd    bool // 该节点是否为字符串的结尾
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a' // ch 变为 node.children 的索引
		// 如果当前指向的子节点不存在
		if node.children[ch] == nil {
			// 创建新的子节点
			node.children[ch] = &Trie{}
		}
		// 指针指向下一个子节点
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		// 说明不存在此前缀
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
