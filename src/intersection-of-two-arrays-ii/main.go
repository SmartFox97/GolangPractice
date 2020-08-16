package main

import "fmt"

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersect(nums1,nums2))
}


func intersect(nums1 []int, nums2 []int) []int {
	numCheck := make(map[int]int)
	for _ , v := range nums1 {
		if _, ok := numCheck[v]; ok {
			numCheck[v] =  numCheck[v] +1
		}else {
			numCheck[v] = 1
		}
	}
	numsRes := []int{}
	for _ , v :=  range nums2 {
		if  value, ok := numCheck[v] ; ok {
			if  value -1 >= 0 {
				numsRes = append(numsRes , v)
				numCheck[v] = value -1
			}
		}
	}
	return numsRes

}
