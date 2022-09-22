package main

type NumArray struct {
	nums []int
	diff []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{nums: nums}
	numArray.diff = make([]int, len(nums))
	numArray.diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		numArray.diff[i] = nums[i] - nums[i-1]
	}
	return numArray
}

// 在闭区间内增加 val，val 可以为负数
func (this *NumArray) increment(i, j int, val int) {
	this.diff[i] += val
	if j+1 < len(this.nums) {
		this.diff[j+1] -= val
	}
}

// 根据差分数组反推原数组
func (this *NumArray) result() []int {
	res := make([]int, len(this.nums))
	res[0] = this.diff[0]
	for i := 1; i < len(this.diff); i++ {
		res[i] = this.diff[i] + res[i-1]
	}
	return res
}
