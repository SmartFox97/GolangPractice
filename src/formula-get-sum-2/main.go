package main

import "fmt"

func getSum(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum1 := 0
		for j := 1; j <= i; j++{
			sum1 += j
		}
		sum = sum*sum1
	}
	return sum
}

func main(){
	for i := 1; i <= 5; i++ {
		fmt.Println("n=",i,"，",getSum(i))
	}
}