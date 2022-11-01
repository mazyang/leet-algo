package main

import "fmt"

// 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 从后向前 寻找两个相邻的升序的元素对 nums[i] < nums[i+1]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 此时 nums[i+1, n-1] 都是降序排列的，然后从 nums[n-1] 开始倒序查找第一个满足大于 nums[i]的较小的大数
	if i >= 0 {
		// 从最后一个元素向前查找
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		// 找到大数后和小数交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// nums[i] 后面的元素变为升序排列
	for x, y := i+1, n-1; x < y; x, y = x+1, y-1 {
		nums[x], nums[y] = nums[y], nums[x]
	}
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func main() {
	a := []int{1, 3, 2}
	nextPermutation(a)
	fmt.Println(a)
}
