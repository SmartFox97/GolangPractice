# 0x00：前言

+ 实现一个函数,参数为n,返回1*(1+2)*…*(1+2+…+n)的值
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4072/)

----------

# 0x01：代码

```golang
package main

import "fmt"

func getSum(n int) (sum int) {
    sum := 1
    for i := 1; i <= n; i++ {
        sum1 := 0
        for j := 1; j <= i; j++{
            sum1 += j
        }
        sum = sum*sum1
    }
    return
}
func main(){
    for i := 1; i <= 5; i++ {
        fmt.Printf("n=%d,result:%d",i,getSum(i))
    }
}
```

----------

# 0x02：运行结果

![result](http://oss.smartfox.cc/2020/06/20/eca72ac081c76.png)
