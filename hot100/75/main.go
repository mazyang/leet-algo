package main

import "fmt"

// 冒泡排序
func sortColors(nums []int) {
	// 每次都将一个最大的元素放到合适的位置
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 方法二：单指针，先将0放到前面，再将1放到0的后面
func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}

func sortColors1(nums []int) {
	count0 := swapColors(nums, 0) // 把 0 排到前面
	swapColors(nums[count0:], 1)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
}

// 方法三：双指针
func sortColors2(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	p, q := 0, n-1
	// q后面的都已经变为2，p前面的已经变为0
	for i := 0; i <= q; i++ {
		// 遇到0放到前面，遇到2放到后面，1不做处理
		if nums[i] == 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
		if nums[i] == 2 {
			nums[i], nums[q] = nums[q], nums[i]
			q--
			// 不是1，可能是0或2，这两个都需要继续交换，所以i需要回退，继续处理
			if nums[i] != 1 {
				i--
			}
		}
	}
}

func main() {
	nums := []int{1, 2, 0}
	sortColors2(nums)
	fmt.Println(nums)
}
