package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

// 方法一：回溯
func readBinaryWatch(turnedOn int) []string {
	hours := []int{1, 2, 4, 8}
	minutes := []int{1, 2, 4, 8, 16, 32}
	var res []string
	path := make([]int, 2)
	var dfs func(int, int, int)
	dfs = func(rest int, hourStart int, minuteStart int) {
		if path[0] > 11 || path[1] > 59 {
			return
		}
		if rest == 0 {
			if path[0] < 12 && path[1] < 60 {
				var tmp string
				tmp = strconv.Itoa(path[0])
				if path[1] < 10 {
					tmp = tmp + ":0" + strconv.Itoa(path[1])
				} else {
					tmp = tmp + ":" + strconv.Itoa(path[1])
				}
				res = append(res, tmp)
			}
			return
		}
		// 选择在小时还是在分钟上加
		var i int
		for i = minuteStart; i < len(minutes); i++ {
			path[1] += minutes[i]
			dfs(rest-1, hourStart, i+1)
			path[1] -= minutes[i]
		}
		var j int
		for j = hourStart; j < len(hours); j++ {
			path[0] += hours[j]
			dfs(rest-1, j+1, i)
			path[0] -= hours[j]
		}
	}
	dfs(turnedOn, 0, 0)
	return res
}

// 方法二：枚举
func readBinaryWatch2(turnedOn int) []string {
	var ans []string
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			// OnesCount8 返回参数的二进制表示形式中 1 的个数
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(readBinaryWatch(2))
}
