package main

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 奇数：二进制表示中，奇数一定比前面那个偶数多一个 1，因为多的就是最低位的 1
		if i%2 == 1 {
			res[i] = res[i-1] + 1
		} else { // 偶数：二进制表示中，偶数中 1 的个数一定和除以 2 之后的那个数一样多
			res[i] = res[i/2]
		}
	}
	return res
}
