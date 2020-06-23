# 0x00：前言

+ 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
+ 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

+ 来源：[力扣（LeetCode）](https://leetcode-cn.com/problems/two-sum)
+ Blog: [噔噔咚](https://www.smartfox.cc/archives/4082/)

+ **Sample Input**

```shell
nums = [2, 7, 11, 15], target = 9
```

+ **Sample Output**

```shell
[0, 1]
```
----------

# 0x01：思路

1. 生成一个map存放值
2. 遍历一次nums中的值，用target与nums中每个元素值相减得到的差在map中进行查找
3. 若map中存在该值，则返回map的value和当前nums元素下标
4. 若map中不存在该值，则将当前nums元素作为key、当前nums元素下标作为value存入map中
5. 若遍历nums完成后，找不到匹配的俩个元素，则返回nil

----------

# 0x02：代码

```golang
package main

import "fmt"

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(nums,target))
}

func twoSum(nums []int, target int) []int {
    keyMap := make(map[int]int,len(nums))
    for k,v := range nums{
        index,ok := keyMap[target - v]
        if ok{
            return []int{index,k}
        }
        keyMap[v] = k
    }
    return nil
}
```

----------

# 0x03：运行结果
![result](https://oss.smartfox.cc/2020/06/21/93a28a2747e6c.png)