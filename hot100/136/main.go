package main

// 只出现一次的数字
func singleNumber(nums []int) int {
	res := 0
	// 使用异或(两个位相同为0 不同为1)，将出现两次的数字消掉，剩下的就是出现一次的数字
	for _, num := range nums {
		res ^= num
	}
	return res
}
