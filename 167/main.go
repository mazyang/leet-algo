package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if target == sum {
			return []int{left, right}
		} else if target > sum {
			left++
		} else if target < sum {
			right++
		}
	}
	return []int{-1, -1}
}
