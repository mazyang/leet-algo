package UnionFind

// UF 并查集
type UF struct {
	// 连通分量的个数
	count int
	// 节点x的父节点为 parent[x]
	parent []int
}

// Constructor n是节点总数
func Constructor(n int) *UF {
	uf := &UF{}
	uf.count = n
	uf.parent = make([]int, n)
	// 父节点指针初始指向自己
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

// Find 返回某个节点x的根节点, 在寻找的同时将所有的树枝拉平, 这样 find 的时间复杂度为 O(1)
func (uf *UF) Find(x int) int {
	// 默认 uf.parent[x] = x, 根节点指向自己
	if uf.parent[x] != x {
		// 路径压缩
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union 将节点q,p连接到一起
func (uf *UF) Union(p, q int) {
	// 找到p,q的根节点
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	// 将两棵树合并为一棵
	uf.parent[rootQ] = rootP
	// 两个分量合二为一
	uf.count--
}

// Connected 判断p,q是否连通
func (uf *UF) Connected(p, q int) bool {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	return rootQ == rootP
}

func (uf *UF) Count() int {
	return uf.count
}
