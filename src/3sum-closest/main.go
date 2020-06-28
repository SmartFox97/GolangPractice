package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1,2,1,-4}
	target := 1
	fmt.Printf("nums=%v \ntarget=%d \nresult=%d\n",nums,target,threeSumClosest(nums,target))
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

