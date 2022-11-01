package main

import "sort"

// 方法一：记录每个字母出现的次数，作为key
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	// [26]byte 作为key
	hashMap := make(map[[26]byte][]string)
	for i := 0; i < len(strs); i++ {
		count := [26]byte{}
		for _, ch := range strs[i] {
			count[ch-'a']++
		}
		hashMap[count] = append(hashMap[count], strs[i])
	}
	for _, value := range hashMap {
		res = append(res, value)
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	var res [][]string
	hashMap := make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		sortedStr := string(tmp)
		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}
	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}

func main() {
	groupAnagrams([]string{"abc"})
}
