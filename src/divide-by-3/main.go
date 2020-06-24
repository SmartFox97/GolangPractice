package main

import "fmt"

func main() {
	var l,r int
	fmt.Scan(&l,&r)
	count:=0
	for i := l; i <= r; i++ {
		if sum(i)%3 == 0{
			count ++
		}
	}
	println(count)

}
func sum(i int) int {
	sum:=0
	for i > 0 {
		sum = sum + i
		i--
	}
	return sum
}

