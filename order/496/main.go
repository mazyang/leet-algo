package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	for index, num := range nums2 {
		hashMap[num] = index
	}
	n := len(nums2)
	greater := make([]int, n)
	var s []int
	for i := n - 1; i >= 0; i-- {
		// 如果当前元素后面的元素小于等于当前元素，出栈
		for len(s) != 0 && s[len(s)-1] <= nums2[i] {
			s = s[:len(s)-1]
		}
		// 栈为空，不满足条件
		if len(s) == 0 {
			greater[i] = -1
		} else {
			// 单调递减栈，最后一个元素就是当前元素的下一个更大的元素
			greater[i] = s[len(s)-1]
		}
		s = append(s, nums2[i])
	}
	var ans []int
	for _, num := range nums1 {
		i := hashMap[num]
		ans = append(ans, greater[i])
	}
	return ans
}
