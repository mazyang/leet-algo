package main

import "leetcode-algo/UnionFind"

// 无向图中连通分量的数目
func countComponents(n int, edges [][]int) int {
	uf := UnionFind.Constructor(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	return uf.Count()
}
