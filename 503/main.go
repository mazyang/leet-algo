package main

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	greater := make([]int, n)
	var s []int
	for i := 2*n - 1; i >= 0; i-- {
		// 如果当前元素后面的元素小于等于当前元素，出栈
		for len(s) != 0 && s[len(s)-1] <= nums[i%n] {
			s = s[:len(s)-1]
		}
		// 栈为空，不满足条件
		if len(s) == 0 {
			greater[i%n] = -1
		} else {
			greater[i%n] = s[len(s)-1]
		}
		s = append(s, nums[i%n])
	}
	return greater
}

func main() {
	nums := []int{1, 2, 3, 4, 3}
	a := nextGreaterElements(nums)
	for i := 0; i < len(a); i++ {
		print(a[i], ' ')
	}
}
