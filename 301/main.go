package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	// 获得待删除的左右括号的数量
	ldel, rdel := 0, 0
	for _, ch := range s {
		if ch == '(' {
			ldel++
		} else if ch == ')' {
			if ldel == 0 {
				rdel++
			} else {
				ldel--
			}
		}
	}
	tdel := ldel + rdel
	var ans []string
	var path []byte
	set := make(map[string]struct{})
	// 删除足够的数量的左右括号
	var dfs func(int, int, int, int, int)
	// lcnt 是path中左括号的数量，rcnt 是path中右括号的数量
	dfs = func(index int, lcnt int, rcnt int, ldel int, rdel int) {
		// 如果删除的时候不满足条件，返回
		if ldel*rdel < 0 || lcnt < rcnt || ldel+rdel > len(s)-index {
			return
		}
		if ldel == 0 && rdel == 0 {
			if len(s)-len(path) == tdel {
				tmp := string(append([]byte{}, path...))
				set[tmp] = struct{}{}
				return
			}
		}
		if index == len(s) {
			return
		}
		ch := s[index]
		if ch == '(' {
			// 选择删除
			dfs(index+1, lcnt, rcnt, ldel-1, rdel)
			// 回溯到此，选择不删除, 如果继续回溯，那么需要撤销选择
			path = append(path, ch)
			dfs(index+1, lcnt+1, rcnt, ldel, rdel)
		} else if ch == ')' {
			dfs(index+1, lcnt, rcnt, ldel, rdel-1)
			path = append(path, ch)
			dfs(index+1, lcnt, rcnt+1, ldel, rdel)
		} else {
			path = append(path, ch)
			dfs(index+1, lcnt, rcnt, ldel, rdel)
		}
		// 撤销选择
		path = path[:len(path)-1]
	}
	dfs(0, 0, 0, ldel, rdel)
	for value := range set {
		ans = append(ans, value)
	}
	return ans
}

func main() {
	s := "()())()"
	fmt.Println(removeInvalidParentheses(s))
}
