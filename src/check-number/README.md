# 0x00：前言

+ 假设有两个未知整数x、y，它们相加等于n，相乘等于m，-10000<n,m<10000 .
+ 请实现一个函数，输入n、m，如果存在x和y则返回True，不存在x和y则返回False.
+ 例如n=5、m=4，则x、y分别为1和4，返回True.
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4078/)

----------

# 0x01：代码

```golang
package main

import (
    "fmt"
    "math"
)

func checkNumber(n, m int) bool {
    res1 := (float64(n)+math.Sqrt(float64(n*n - 4*m)))/2
    res2 := (float64(n)-math.Sqrt(float64(n*n - 4*m)))/2
    if res1 == float64(int(res1)) || res2 == float64(int(res2)) {
        return true
    }
    return false
}

func main() {
    n := 5
    m := 4
    fmt.Printf("n = %d , m = %d ,CheckStatus: %v",n,m,checkNumber(5,4))
}

```

----------

# 0x02：运行结果

![result](https://oss.smartfox.cc/2020/06/20/2714c468b32e0.png)