# 0x00：前言

+ 实现一个函数，参数为字符串a和b，字符串是分数格式，如1/4，返回a+b的和的分数格式字符串。
+ 例如：参数为11/8和3/4时，返回17/8
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4076/)

----------

# 0x01：代码

```golang
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func getSum(a,b string)(result string){
    arr1 := strings.Split(a,"/")
    arr2 := strings.Split(b,"/")
    i, _ := strconv.Atoi(arr1[0])
    j, _ := strconv.Atoi(arr1[1])
    k, _ := strconv.Atoi(arr2[0])
    l, _ := strconv.Atoi(arr2[1])
    denominator := getLowestCommonMultiple(j,l)
    numerator := (i*getLowestCommonMultiple(j,l)/j)+(k*getLowestCommonMultiple(j,l)/l)
    gcd := getGreatestCommonDivisor(numerator,denominator)
    result = fmt.Sprintf("%d/%d",numerator/gcd,denominator/gcd)
    return
}


func getLowestCommonMultiple(a,b int) int{
    if a<b{
        a,b = b,a
    }
    i:=1
    for i<=b{
        if a*i%b == 0{
            return a*i
        }
        i++
    }
    return a*i
}

func getGreatestCommonDivisor(a, b int) int {
    if a < b {
        a ,b = b ,a
    }
    if a%b == 0 {
        return b
    }
    return getGreatestCommonDivisor(a%b,b)
}
func main() {
    a := "11/8"
    b := "3/4"
    fmt.Println(getSum(a,b))
}
```

----------

# 0x02：运行结果

![result](https://oss.smartfox.cc/2020/06/20/6f42098c0e2e0.png)