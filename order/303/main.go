package main

type NumArray struct {
	nums []int
	sum  []int
}

func Constructor(nums []int) NumArray {
	var numArray NumArray
	numArray.nums = nums
	// sum 保存nums[0]到nums[i-1] 的累加和
	sum := make([]int, len(nums)+1)
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	numArray.sum = sum
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}
