package main

import (
	"math/rand"
)

type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:       make([]int, 0),
		valToIndex: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	// 如果已经存在
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	this.valToIndex[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// 如果不存在
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	// 获得被删除元素的索引
	index := this.valToIndex[val]
	// 修改最后一个元素的索引
	this.valToIndex[this.nums[len(this.nums)-1]] = index
	// 元素交换
	this.nums[index], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[index]
	// 删除最后一个元素
	this.nums = this.nums[:len(this.nums)-1]
	// 删除val对应的索引
	delete(this.valToIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
