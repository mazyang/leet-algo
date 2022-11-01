package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	left, right := (m+n+1)/2, (m+n+2)/2
	var findKth func(i, j, k int) int
	// i,j分别代表在nums1和nums2中的位置
	// k代表在nums1和nums2剩余的元素中查找第k个元素(k从1开始)，也就相当于分别在nums1和nums2中查找第k/2个元素
	findKth = func(i, j, k int) int {
		// nums1被淘汰，在nums2中查找第k个
		if i >= m {
			return nums2[j+k-1]
		}
		// nums2被淘汰，在nums1中查找第k个
		if j >= n {
			return nums1[i+k-1]
		}
		if k == 1 {
			return min(nums1[i], nums2[j])
		}
		midVal1 := math.MaxInt32
		midVal2 := math.MaxInt32
		// 判断在nums1的剩余元素中是否存在k/2个元素
		if i+k/2-1 < m {
			midVal1 = nums1[i+k/2-1]
		}
		// 判断在nums2的剩余元素中是否存在k/2个元素
		if j+k/2-1 < n {
			midVal2 = nums2[j+k/2-1]
		}
		// midVal1比较小，说明要找的数字不在nums1的前k/2个数字，nums1的前k/2个数字被丢弃
		if midVal1 < midVal2 {
			return findKth(i+k/2, j, k-k/2)
		}
		return findKth(i, j+k/2, k-k/2)
	}
	return (float64(findKth(0, 0, left)) + float64(findKth(0, 0, right))) / 2.0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
