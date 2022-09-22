package main

import "fmt"

// debug 技巧

var count int

// 打印缩进, 代表是第几层递归
func printIndent(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("  ")
	}
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// 在递归内部打印
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		// 进入递归时先增加缩进空格数量, 表明当前层级
		count++
		printIndent(count)
		// 打印输入参数
		fmt.Printf("i=%d j=%d \n", i, j)

		if i == -1 {
			// 退出时先打印空格，再减少空格数量，表明退出当前层级
			printIndent(count)
			count--
			// 退出时打印结果
			fmt.Printf("return %d \n", j+1)
			return j + 1
		}
		if j == -1 {
			printIndent(count)
			count--
			fmt.Printf("return %d \n", j+1)
			return i + 1
		}
		res := 0
		// 进行状态转移
		if word1[i] == word2[j] {
			res = dfs(i-1, j-1)
		} else {
			res = min(dfs(i, j-1), min(dfs(i-1, j), dfs(i-1, j-1))) + 1
		}
		printIndent(count)
		count--
		fmt.Printf("return %d \n", res)
		return res
	}
	return dfs(m-1, n-1)
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
