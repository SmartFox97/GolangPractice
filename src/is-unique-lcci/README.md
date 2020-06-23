# 0x00：前言

+ 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
+ 来源：[力扣(LeetCode)](https://leetcode-cn.com/problems/is-unique-lcci)
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4087/)

+ **Sample Input**

```shell
s1 = "leetcode"
s2 = "abc"
```
+ **Sample Output**

```shell
false
true
```

+ 限制：

> 0 <= len(s) <= 100
> 如果你不使用额外的数据结构，会很加分。

----------

# 0x01：思路

+ 简单粗暴用MAP

1. 声明并初始化一个map用于存放值；
2. 将传入值转化为rune（这一步是为了考虑中文字符）；
3. 遍历转化后的字符串，取出字符，放入map中进行查询；
4. 若字符已存在map中，函数返回false；
5. 若字符不存在map中，则将字符存入map进入下一次循环；
6. 若循环结束后，并没有在map中找到该字符，则返回true.

----------

# 0x02：代码

```golang
package main

import "fmt"

func main() {
    s1 := "leetcode"
    s2 := "abc"
    fmt.Printf("Input：%s,result：%v \n",s1,isUnique(s1))
    fmt.Printf("Input：%s,result：%v \n",s2,isUnique(s2))
}

func isUnique(astr string) bool {
    str := []rune(astr)
    strMap := make(map[rune]int)
    for i , ch := range str{
        _,ok := strMap[ch]
        if ok {
            return false
        }
    strMap[ch] = i
    }
return true
}

```

----------

# 0x03：运行结果

![result](https://oss.smartfox.cc/2020/06/22/33fd0ed08e5cf.png)