package main

// 单词拆分
// 方法一：动态规划
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	// 将单词放入哈希表中
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	// dp[i] 代表 s[0...i-1] 字符串是否可以被拼接成功
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 如果 s[0...j-1] 可以被拼接，而且存在单词 s[j:i]，说明 s[0...i-1] 也可以被拼接
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
