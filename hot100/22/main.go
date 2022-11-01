package main

// 括号生成
// 条件一：对于一个合法的括号组合，左右括号的数量一定相等
// 条件二：对于一个合法的括号组合，在任意位置p，[0:p] 中左括号的数量一定大于右括号的数量
func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	var recur func(int, int)
	// 剩余左右括号的数量
	recur = func(left int, right int) {
		// base case
		if left == 0 && right == 0 {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		if left < 0 || right < 0 {
			return
		}
		// 剩余的左括号数量多，不满足条件二
		if left > right {
			return
		}
		// 尝试添加左括号
		path = append(path, '(')
		recur(left-1, right)
		path = path[:len(path)-1]
		// 尝试添加右括号
		path = append(path, ')')
		recur(left, right-1)
		path = path[:len(path)-1]
	}
	recur(n, n)
	return res
}
