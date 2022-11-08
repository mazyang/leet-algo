package main

import "fmt"

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	// 前缀和
	arr1 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr1[i] = arr1[i-1] + nums[i-1]
	}
	max := func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}
	var res int64
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i-j+1 == k && !judge(nums, j, i) {
				res = max(res, int64(arr1[i+1]-arr1[j]))
			}
		}
	}
	return res
}

func judge(nums []int, start, end int) bool {
	hashMap := make(map[int]int)
	for i := start; i <= end; i++ {
		if _, ok := hashMap[nums[i]]; ok {
			return true
		}
		hashMap[nums[i]]++
	}
	return false
}

func main() {
	nums := []int{5, 1, 3}
	fmt.Println(maximumSubarraySum(nums, 1))
}
