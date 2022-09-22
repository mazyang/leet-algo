package main

func plusOne(s string, i int) string {
	ch := []byte(s)
	if ch[i] == '9' {
		ch[i] = '0'
	} else {
		ch[i] += 1
	}
	return string(ch)
}

func minusOne(s string, i int) string {
	ch := []byte(s)
	if ch[i] == '0' {
		ch[i] = '9'
	} else {
		ch[i] -= 1
	}
	return string(ch)
}

func openLock(deadends []string, target string) int {
	res := 0
	deads := make(map[string]struct{})
	for i := 0; i < len(deadends); i++ {
		deads[deadends[i]] = struct{}{}
	}
	visited := make(map[string]struct{})
	visited["0000"] = struct{}{}
	queue := []string{"0000"}
	for len(queue) != 0 {
		for i := len(queue) - 1; i >= 0; i-- {
			cur := queue[0]
			queue = queue[1:]
			if _, ok := deads[cur]; ok {
				continue
			}
			if cur == target {
				return res
			}
			// 共有四个位置，每个位置有+1 -1 两种选择
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				// 如果没有被访问过
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = struct{}{}
				}
				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = struct{}{}
				}
			}
		}
		res++
	}
	return -1
}

func main() {
	openLock([]string{"8888"}, "0009")
}
