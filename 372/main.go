package main

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	part1 := myPow(a, last)
	part2 := myPow(superPow(a, b), 10)
	return (part1 * part2) % 1337
}

// 可能会溢出
func myPow(a, k int) int {
	res := 1
	for i := 0; i < k; i++ {
		res *= a
		res %= 1337
	}
	return res
}

// 递归方法
func myPow2(a, k int) int {
	if k == 0 {
		return 1
	}
	if k%2 == 1 {
		return (a * myPow2(a, k-1)) % 1337
	} else {
		sub := myPow2(a, k/2)
		return (sub * sub) % 1337
	}
}

func main() {
	var a int
	a = 932 * 2147483647
	println(a)
}
