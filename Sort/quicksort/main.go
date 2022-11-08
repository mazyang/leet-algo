package main

import (
	"fmt"
	"math/rand"
)

// 快速排序
func sort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(nums, 0, len(nums)-1)
	sort(nums, left, p-1)
	sort(nums, p+1, right)
}

func partition(nums []int, left, right int) int {
	// 也可以先洗牌
	randomIdx := left + rand.Intn(right-left+1)
	// 交换第一个元素和随机位置的元素
	nums[left], nums[randomIdx] = nums[randomIdx], nums[left]
	// 获得分界点
	pivot := nums[left]
	i, j := left+1, right
	for i <= j {
		// 从左到右寻找第一个大于分界点的元素
		for i <= j && nums[i] <= pivot {
			i++
		}
		// 从右到左寻找第一个小于分界点的元素
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		// 找到之后交换，继续寻找
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	sort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
