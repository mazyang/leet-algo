package main

// 方法一：使用额外的空间
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	count := make([]int, n+1)
	for i := 0; i < n; i++ {
		count[nums[i]]++
	}
	var res []int
	for i := 1; i <= n; i++ {
		if count[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

// 方法二：映射
func findDisappearedNumbers2(nums []int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for _, num := range nums {
		// 如果对应索引上的数字存在，变为负数
		if nums[abs(num)-1] > 0 {
			nums[abs(num)-1] *= -1
		}
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
