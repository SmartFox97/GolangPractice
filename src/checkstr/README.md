# 0x00：前言

+ 给出一串字符串，求最少需要删除多少个字符，才能使得剩下的字符串满足以下性质：
+ 出现最多的字符的数量大于等于字符串长度的一半

+ 提示：
> + 输入包含一行，为给出的字符串。（字符串长度小于等于100000，保证全是小写字母）

+ **Sample Input**
```shell script
abcdaabdda
```

+ **Sample Output**
```shell script
2
```

+ 来源：[牛客网](https://www.nowcoder.com/)

----------

# 0x01：思路

TODO

----------

# 0x02：代码

```golang
package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(checkstr(a))
}

func checkstr(str string) int {
	strlen := len(str)
	max := 0
	for i := 'a'; i <= 'z'; i++ {
		a:= string(i)
		n:= strings.Count(str,a)
		if max < n{
			max = n
		}

	}
	del := strlen - max * 2
	return del

}

```

----------

# 0x02：运行结果

![result](http://oss.smartfox.cc/2020/06/24/877d3fc18cffe.png)