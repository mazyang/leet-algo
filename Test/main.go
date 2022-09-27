package main

// 单调递增栈：当前元素大于栈顶元素就一直出栈直到满足条件(当前元素大于栈顶元素或栈为空)
func nextGreaterElement(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	var s []int
	for i := n - 1; i >= 0; i-- {
		// 如果当前元素后面的元素小于等于当前元素，出栈
		for len(s) != 0 && s[len(s)-1] <= nums[i] {
			s = s[:len(s)-1]
		}
		// 栈为空，不满足条件
		if len(s) == 0 {
			res[i] = -1
		} else {
			res[i] = s[len(s)-1]
		}
		s = append(s, nums[i])
	}
	return res
}

func main() {
	nums := []int{2, 1, 2, 4, 3}
	a := nextGreaterElement(nums)
	for _, i := range a {
		print(i, " ")
	}
}
