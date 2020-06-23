# 0x00：前言

+ 给定一条公式“1-3-5+7-9-11+...+n”
+ 请设计一个函数，假设传入一个奇数值n，返回这条公式的计算结果
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4083/)


+ **Sample Input**

```shell
n = 1
n = 5
n = 11
n = 99
```

+ **Sample Output**

```shell
1
-7
-20
-834
```
----------

# 0x01：思路

+ 暴力思路

1. 声明标记 j ，默认值为0；声明一个sum存放计算结果；声明一个i存放相加值；
2. 当 j 为0时，sum = sum + i，i = i + 2，进入下一次循环；
3. 当 j 为1时，sum = sum - i，i = i + 2，进入下一次循环；
4. 当 j 为2时，sum = sum - i，i = i + 2，且 j 的数值重置为0，进入下一次循环；
5. 循环结束返回sum的值

----------

# 0x02：代码

```golang
package main

import "fmt"

func main() {
    fmt.Println(getSum(1))
    fmt.Println(getSum(5))
    fmt.Println(getSum(11))
    fmt.Println(getSum(99))
}

func getSum(n int) int{
    sum := 0
    i := 1
    j := 0
    for i <= n {
        if j == 1 {
            sum = sum - i
            i = i + 2
            j++
            continue
        }
        if j == 2 {
            sum = sum - i
            i = i + 2
            j = 0
            continue
        }
        sum = sum + i
        j++
        i = i + 2
    }
    return sum
}

```

----------

# 0x03：运行结果

![result](http://oss.smartfox.cc/2020/06/22/e7d7355bfbffb.png)