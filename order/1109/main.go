package main

// 方法：差分数组
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	diff := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		// 在范围bookings[i][0]到bookings[i][1] 之间加上bookings[i][2]
		diff[bookings[i][0]-1] += bookings[i][2]
		if bookings[i][1] < n {
			diff[bookings[i][1]] -= bookings[i][2]
		}
	}
	res[0] = diff[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}
