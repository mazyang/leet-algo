package main

import (
	"fmt"
	"math"
)

func splitIntoFibonacci(num string) []int {
	var F []int
	n := len(num)
	var backtrack func(index, sum, prev int) bool
	// sum 为最后两个数的和, index 是下一个数的起始位置
	backtrack = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}
		cur := 0 // cur 为当前切割的数字
		for i := index; i < n; i++ {
			// 每个块的数字一定不要以零开头，除非这个块是数字 0 本身
			if i > index && num[index] == '0' {
				break
			}
			// int 类型在64位操作系统中是 int64
			cur = cur*10 + int(num[i]-'0')
			// 拆出的整数要符合 32 位有符号整数类型
			if cur > math.MaxInt32 {
				break
			}
			// F[i] + F[i+1] = F[i+2]
			if len(F) >= 2 {
				// 如果当前切割的值小于前两个数之和，继续增加
				if cur < sum {
					continue
				}
				// 大于则剪枝
				if cur > sum {
					break
				}
			}
			// cur 符合要求，加入序列 F
			F = append(F, cur)
			if backtrack(i+1, prev+cur, cur) {
				return true
			}
			F = F[:len(F)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return F
}

func main() {
	num := "1101111"
	fmt.Println(splitIntoFibonacci(num))
}
