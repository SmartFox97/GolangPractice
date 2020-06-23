package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5
	m := 4
	fmt.Printf("n = %d , m = %d ,CheckStatus: %v",n,m,checkNumber(5,4))
}

func checkNumber(n, m int) bool {
	res1 := (float64(n)+math.Sqrt(float64(n*n - 4*m)))/2
	res2 := (float64(n)-math.Sqrt(float64(n*n - 4*m)))/2
	if res1 == float64(int(res1)) || res2 == float64(int(res2)) {
		return true
	}
	return false
}
