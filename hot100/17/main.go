package main

// 回溯
func letterCombinations(digits string) []string {
	hashMap := make(map[byte]string)
	hashMap['2'] = "abc"
	hashMap['3'] = "def"
	hashMap['4'] = "ghi"
	hashMap['5'] = "jkl"
	hashMap['6'] = "mno"
	hashMap['7'] = "pqrs"
	hashMap['8'] = "tuv"
	hashMap['9'] = "wxyz"
	var res []string
	var path []byte
	var recur func(string)
	recur = func(digits string) {
		if len(digits) == 0 {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		alpha := hashMap[digits[0]]
		for _, val := range alpha {
			path = append(path, byte(val))
			recur(digits[1:])
			path = path[:len(path)-1]
		}
	}
	recur(digits)
	return res
}
