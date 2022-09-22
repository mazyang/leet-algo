package main

import (
	"fmt"
	"leetcode-algo/UnionFind"
)

func equationsPossible(equations []string) bool {
	uf := UnionFind.Constructor(26)
	// == 两边的字母建立连通性
	for _, eq := range equations {
		if eq[1] == '=' {
			x := eq[0] - 'a'
			y := eq[3] - 'a'
			uf.Union(int(x), int(y))
		}
	}
	// 检查 != 关系是否打破了已经建立的连通性
	for _, eq := range equations {
		if eq[1] == '!' {
			x := eq[0] - 'a'
			y := eq[3] - 'a'
			// 破坏了连通性, 返回false
			if uf.Connected(int(x), int(y)) {
				return false
			}
		}
	}
	return true
}

func main() {
	eq := []string{"a==b", "b!=a"}
	fmt.Println(equationsPossible(eq))
}
