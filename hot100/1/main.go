package main

// 使用哈希表
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for index, num := range nums {
		if p, ok := hashMap[target-num]; ok {
			return []int{index, p}
		}
		hashMap[num] = index
	}
	return []int{-1, -1}
}
