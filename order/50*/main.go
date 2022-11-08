package main

// 快速幂
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	// 如果幂小于0，那么将底数变为倒数，获得幂的绝对值
	res := 1.0
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n > 0 {
		// 等于 n % 2 == 1，相当于判断 n 的二进制的最后一位是否为1
		if n&1 == 1 {
			// 如果 n=11=1011(二进制) = x^1 * x^2 * x^8
			res *= x
		}
		// x^1  x^2  x^4  x^8  x^16 ......
		x = x * x
		// n 右移一位
		n = n >> 1
	}
	return res
}
