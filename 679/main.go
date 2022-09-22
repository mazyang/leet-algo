package main

const (
	TARGET                          = 24
	EPSILON                         = 1e-6 // 误差
	ADD, MULTIPLY, SUBTRACT, DIVIDE = 0, 1, 2, 3
)

func judgePoint24(nums []int) bool {
	list := []float64{}
	// 将原数字转为float64
	for _, num := range nums {
		list = append(list, float64(num))
	}
	return solve(list)
}

func solve(list []float64) bool {
	if len(list) == 0 {
		return false
	}
	if len(list) == 1 {
		return abs(list[0]-TARGET) < EPSILON
	}
	size := len(list)
	// 从list中挑选一个数
	for i := 0; i < size; i++ {
		// 从list中再挑选一个数
		for j := 0; j < size; j++ {
			// 挑选的两个数不能相同
			if i != j {
				var list2 []float64
				// 将未被挑选的数加入list2中
				for k := 0; k < size; k++ {
					if k != i && k != j {
						list2 = append(list2, list[k])
					}
				}
				for k := 0; k < 4; k++ {
					// k < 2 是指挑选的运算是 + 和 *, 这两个运算满足交换律
					// i < j 是为了满足交换律的成立，减少计算时间
					//eg: i = 0, j = 1, k = 1 → 挑第一个和第二个数，两数相乘
					//    i = 1, j = 0, k = 1 → 挑第二个和第一个数，两数相乘  →  无效的重复计算
					if k < 2 && i < j {
						continue
					}
					switch k {
					case ADD:
						list2 = append(list2, list[i]+list[j])
					case MULTIPLY:
						list2 = append(list2, list[i]*list[j])
					case SUBTRACT:
						list2 = append(list2, list[i]-list[j])
					case DIVIDE:
						//如果除数等于0，直接判这次所挑选的运算为false，跳出此次循环
						if abs(list[j]) < EPSILON {
							continue
						} else {
							list2 = append(list2, list[i]/list[j])
						}
					}
					if solve(list2) {
						return true
					}
					list2 = list2[:len(list2)-1]
				}
			}
		}
	}
	return false
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
