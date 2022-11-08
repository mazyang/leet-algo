package main

func minInsertions(s string) int {
	res, need := 0, 0 // need为需要增加的右括号的数量，res是需要增加的左右括号的数量
	for _, c := range s {
		if c == '(' {
			// 出现一个左括号，就需要两个右括号对应
			need += 2
			// 如果需要的右括号的数量不为偶数
			if need%2 == 1 {
				// 插入一个右括号
				res++
				need--
			}
		}
		if c == ')' {
			need--
			// 右括号数量多了一个
			if need == -1 {
				// 增加一个左括号
				res++
				// 增加一个左括号就要有两个右括号对应，现在右括号多一个，抵消后再增加一个就可以了
				need = 1
			}
		}
	}
	return res + need
}
