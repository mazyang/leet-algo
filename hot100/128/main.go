package main

func longestConsecutive(nums []int) int {
	// 使用hash作为集合, numSet中没有重复元素
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	res := 0
	for num := range numSet {
		// 判断上一个是否在集合中，在集合中就跳过，不在集合中就执行下面的代码
		// 要枚举的数 x 一定是在数组中不存在前驱数 x−1 的，这样可以避免重复枚举
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			// 当前数字的下一个是否还在集合中，如果存在就更新
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if res < currentStreak {
				res = currentStreak
			}
		}
	}
	return res
}
