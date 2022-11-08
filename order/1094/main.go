package main

// 差分数组：原数组保存每站车上的人数
func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	res := make([]int, 1001)
	for i := 0; i < len(trips); i++ {
		val := trips[i][0]
		diff[trips[i][1]] += val
		if trips[i][2] < 1001 {
			// 乘客在车上的区间是 [trip[i][1]，trip[i][2])，是一个左闭右开的区间
			diff[trips[i][2]+1-1] -= val
		}
	}
	res[0] = diff[0]
	for i := 1; i < 1001; i++ {
		res[i] = diff[i] + res[i-1]
	}
	for i := 0; i < 1001; i++ {
		if capacity < res[i] {
			return false
		}
	}
	return true
}
