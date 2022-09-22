package main

func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	var backtrack func(int, int)
	// left 和 right 是可用的左右括号数目
	backtrack = func(left int, right int) {
		if left > right {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		// 尝试放一个左括号
		path = append(path, '(')
		backtrack(left-1, right)
		path = path[:len(path)-1]

		// 尝试放一个右括号
		path = append(path, ')')
		backtrack(left, right-1)
		path = path[:len(path)-1]
	}
	backtrack(n, n)
	return res
}
