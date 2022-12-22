package main

import "fmt"

// 将一个数组抽象为一个环形链表，然后按照环形链表寻找环入口的方式寻找中重复元素
// 对 nums 数组建图，每个位置 i 连一条 i→nums[i] 的边
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {

	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	//findDuplicate([]int{1, 3, 4, 2, 2})
	fmt.Println(255 ^ 3)
}
