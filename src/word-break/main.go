package main

import (
	"fmt"
)

func main() {
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s1,wordDict1))
	fmt.Println(wordBreak(s2,wordDict2))
	fmt.Println(wordBreak(s3,wordDict3))
}

func wordBreak(s string, wordDict []string) bool {
	strlen := len(s)
	dp := make([]bool,strlen + 1)
	wordMap := make(map[string]bool)
	for _,v := range wordDict{
		wordMap[v] = true
	}
	dp[0] = true

	wordMap[""] = true
	for i := 0; i <= strlen; i++ {
		for j := 0; j <= i; j++ {
			_,ok := wordMap[s[j:i]]
			if ok && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[strlen]
}