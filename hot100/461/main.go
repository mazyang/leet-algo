package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {
	tmp := x ^ y
	return bits.OnesCount(uint(tmp))
}

func main() {
	fmt.Println(hammingDistance(1, 3))
}
