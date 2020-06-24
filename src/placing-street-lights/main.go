package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scan(&t)
	var counts []int
	for i := 0; i < t; i++ {
		n:=0
		var str string
		fmt.Scan(&n)
		fmt.Scan(&str)
		road := strings.Split(str,"")
		j,count := 0,0
		for j < n {
			if road[j] == "." {
				count++
				j+=3
			}else {
				j++
			}
		}
		counts = append(counts,count)
	}
	for _,v := range counts {
		fmt.Println(v)
	}
}
