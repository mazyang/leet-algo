package main

func twoSum(nums []int, target int) []int {
	map1 := map[int]int{}
	// 每次遍历都在当前数字前面的所有数字中搜索，是否满足条件，然后再加入hashMap中
	for i, num := range nums {
		if p, ok := map1[target-num]; ok {
			return []int{p, i}
		}
		map1[num] = i
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for idx, val := range nums {
		hashMap[val] = idx
	}
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		j, ok := hashMap[other]
		if ok && j != i {
			return []int{i, j}
		}
	}
	return []int{-1, -1}
}
