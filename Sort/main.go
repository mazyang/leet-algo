package main

import (
	"fmt"
	"sort"
)

// 方法一

type MyIntList []int

func (x MyIntList) Len() int {
	return len(x)
}

func (x MyIntList) Less(i, j int) bool {
	return x[i] > x[j]
}

func (x MyIntList) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	a := []int{1, 2, 3, 4}
	// 方法一
	sort.Sort(MyIntList(a))
	// 方法二
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	// 方法三
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}
