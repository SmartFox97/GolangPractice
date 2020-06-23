package main

import "fmt"

func main() {
	fmt.Println(getSum(1))
	fmt.Println(getSum(5))
	fmt.Println(getSum(11))
	fmt.Println(getSum(99))
}

func getSum(n int) int{
	sum := 0
	i := 1
	j := 0
	for i <= n {
		if j == 1 {
			sum = sum - i
			i = i + 2
			j++
			continue
		}
		if j == 2 {
			sum = sum - i
			i = i + 2
			j = 0
			continue
		}
		sum = sum + i
		j++
		i = i + 2
	}
	return sum
}
