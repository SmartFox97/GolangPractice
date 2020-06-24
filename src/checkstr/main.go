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
