# 0x00：前言

+ 给你两个二进制字符串，返回它们的和（用二进制表示）
+ 输入为 非空 字符串且只包含数字 1 和 0


+ **Sample Input**

```shell
a = "0", b = "0"
a = "11", b = "1"
a = "1010", b = "1011"
```

+ **Sample Output**

```shell
"0"
"100"
"10101"
```

+ 提示：

> + 每个字符串仅由字符 '0' 或 '1' 组成。
> + 1 <= a.length, b.length <= 10^4
> + 字符串如果不是 "0" ，就都不含前导零。

+ 来源：[力扣（LeetCode）](https://leetcode-cn.com/problems/add-binary)
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4088/)

----------

# 0x01：思路

1. 开局先检查a、b字符串的长度，并让a串始终都是最长的串（即a串长度始终大于b串）；
2. 从字符串末尾开始相加，相加结果大于等于2则需进位；
3. 若长短串，则将长串剩余部分给他接上；
4. 最后无论flag是否为0，都将把flag写进结果串首位；
5. 检查结果串的首位字符是否为0，如果是则从第一个非零字符开始返回，否则直接返回；

----------

# 0x02：代码

```golang
package main

import (
    "fmt"
)

func main() {
    fmt.Println(addBinary("0","0"))
    fmt.Println(addBinary("11","1"))
    fmt.Println(addBinary("1010","1011"))
}

func addBinary(a, b string) string {
    sum := 0
    flag:= 0
    lenA := len(a)
    lenB := len(b)
    if lenA < lenB{
        lenA,lenB = lenB,lenA
        a,b = b,a
    }
    result := make([]byte , lenA + 1)
    for lenB > 0 {
        lenA --
        lenB --
        sum = int(a[lenA] - '0') + int(b[lenB] - '0') + flag
        flag = sum / 2
        sum = sum % 2
        result[lenA + 1] = byte(sum + '0')
    }
    for lenA > 0 {
        lenA --
        sum = int(a[lenA] - '0') + flag
        flag = sum / 2
        sum = sum % 2
        result[lenA + 1] = byte(sum + '0')
    }
    result[0] = byte(flag + '0')

    for lenA < len(result)-1 {
        if result[lenA] == '0' {
            lenA++
        } else {
            break
        }
    }
    return string(result[lenA:])
}
```

----------

# 0x02：运行结果

![result](https://oss.smartfox.cc/2020/06/23/2f83415cf6da3.png)
