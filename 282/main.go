package main

func addOperators(num string, target int) []string {
	var ans []string
	n := len(num)
	var backtrack func(expr []byte, start, res, mul int)
	// expr是当前构造出来的表达式，res是当前表达式的计算结果，mul为表达式最后一个连乘串的计算结果（为了实现乘法优先）
	backtrack = func(expr []byte, start, res, mul int) {
		// 枚举结束
		if start == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		// signIndex是待插入符号的位置
		signIndex := len(expr)
		if start > 0 {
			expr = append(expr, 0) // 占位，下面填充符号
		}
		// 枚举截取的数字长度（取多少位），注意数字可以是单个 0 但不能有前导零
		// start 是当前选择的数字的第一个，所以要么 num[start]!='0'，要么j==start(代表截取的只有一个0)
		for j, val := start, 0; j < n && (j == start || num[start] != '0'); j++ {
			val = val*10 + int(num[j]-'0') // val 是当前截取的值
			expr = append(expr, num[j])
			if start == 0 { // 表达式开头不能添加符号
				backtrack(expr, j+1, val, val)
			} else { // 枚举符号
				expr[signIndex] = '+'
				backtrack(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				backtrack(expr, j+1, res-val, -val)
				expr[signIndex] = '*'
				// 例如：1+2*3，到达此行的时候，res=3, mul=2, val=3，这样就可以实现乘法优先
				backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)
	return ans
}
