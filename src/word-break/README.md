# 0x00：前言

+ 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

+ 说明：
> + 拆分时可以重复使用字典中的单词。
> + 你可以假设字典中没有重复的单词。

+ **Sample Input**

```shell
s1 = "leetcode", wordDict1 = ["leet", "code"]
s2 = "applepenapple", wordDict2 = ["apple", "pen"]
s3 = "catsandog", wordDict3 = ["cats", "dog", "sand", "and", "cat"]
```

+ **Sample Output**

```shell
true   //返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
true   //返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
false
```

+ 来源：[力扣（LeetCode）](https://leetcode-cn.com/problems/word-break)
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4088/)

----------

# 0x01：思路

动态规划
1.
2.
3.
4.

----------

# 0x02：代码

```golang
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
```

----------

# 0x02：运行结果

![result](http://oss.smartfox.cc/2020/06/25/04c1b3d99a015.png)
