package main

import "fmt"

// 共有n个楼梯，每次只能爬1或2阶，计算所有的上楼方式
func f1(n int) [][]int {
	var res [][]int
	var path []int
	var recur func(int)
	recur = func(n int) {
		if n < 0 {
			return
		}
		if n == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 1; i <= 2; i++ {
			path = append(path, i)
			recur(n - i)
			path = path[:len(path)-1]
		}
	}
	recur(n)
	return res
}

func main() {
	fmt.Println(f1(3))
}