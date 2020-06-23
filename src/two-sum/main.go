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
