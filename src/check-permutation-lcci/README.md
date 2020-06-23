# 0x00：前言

+ 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

+ **Sample Input**

```shell
s1 = "abc", s2 = "bca"
s1 = "abc", s2 = "bad"
```

+ **Sample Output**

```shell
true
false
```

+ 说明：

> + 0 <= len(s1) <= 100
> + 0 <= len(s2) <= 100

+ 来源：[力扣（LeetCode）](https://leetcode-cn.com/problems/check-permutation-lcci)
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4089/)

----------

# 0x01：思路

1. 先将传入的s1和s2串转化为字符切片；
2. 使用sort对俩切片进行排序；
3. 再将俩切片转化为字符串；
4. 对比俩字符串是否一致；

----------

# 0x02：代码

```golang
package main

import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    fmt.Println(CheckPermutation("abc", "bca"))
    fmt.Println(CheckPermutation("abc", "bad"))
}

func CheckPermutation(s1 string, s2 string) bool {
    str1 := strings.Split(s1,"")
    str2 := strings.Split(s2,"")
    sort.Strings(str1)
    sort.Strings(str2)
    tmp1 := strings.Join(str1,"")
    tmp2 := strings.Join(str2,"")
    return tmp2 == tmp1
}
```

----------

# 0x02：运行结果

![result](https://oss.smartfox.cc/2020/06/23/7e4992bf3b669.png)
