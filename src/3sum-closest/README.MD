# 0x00：前言

+ 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

+ **Sample Input**
```shell script
nums = [-1,2,1,-4]
target = 1
```

+ **Sample Output**
```shell script
2
```
+ 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

+ 提示：
> + 3 <= nums.length <= 10^3
> + -10^3 <= nums[i] <= 10^3
> + -10^4 <= target <= 10^4

来源：[力扣（LeetCode）](https://leetcode-cn.com/problems/3sum-closest)

----------

# 0x01：思路

双指针法
1. 先将传入的数组nums进行排序；
2. 声明俩个下标n，m；一个从左侧向右移动，一个从右向左移动；
3. 开局先让数组最小的三个数相加并赋给result
4. 进入循环，每次循环都有sums = sum[i]+sum[n]+sum[m]
5. 若result与target的距离比sums与target的距离大，则将sums赋给result
6. 每次循环中，如果sums比target大，则n-1（大值变小）；反之，m-1（小值变大）

----------

# 0x02：代码
```golang
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1,2,1,-4}
	target := 1
	fmt.Println(threeSumClosest(nums,target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		m, n := i+1, len(nums)-1
		for m < n {
			sums := nums[i] + nums[m] + nums[n]
			if math.Abs(float64(sums-target)) < math.Abs(float64(result-target)) {
				result = sums
			}
			if sums > target {
				n--
			} else if sums < target {
				m++
			} else {
				return sums
			}
		}
	}
	return result
}
```

----------

# 0x02：运行结果

![result](http://oss.smartfox.cc/2020/06/24/03ccf6e6b8745.png)