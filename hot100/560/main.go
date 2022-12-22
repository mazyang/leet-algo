package main

func subarraySum(nums []int, k int) int {
	res, pre := 0, 0
	// m[i]=j 代表前缀和为i的个数为j
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			res += m[pre-k]
		}
		m[pre] += 1
	}
	return res
}
