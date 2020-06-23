package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "11/8"
	b := "3/4"
	fmt.Println(getSum(a,b))
}
func getSum(a,b string)(result string){
	arr1 := strings.Split(a,"/")
	arr2 := strings.Split(b,"/")
	i, _ := strconv.Atoi(arr1[0])
	j, _ := strconv.Atoi(arr1[1])
	k, _ := strconv.Atoi(arr2[0])
	l, _ := strconv.Atoi(arr2[1])
	denominator := getLowestCommonMultiple(j,l)
	numerator := (i*getLowestCommonMultiple(j,l)/j)+(k*getLowestCommonMultiple(j,l)/l)
	gcd := getGreatestCommonDivisor(numerator,denominator)
	result = fmt.Sprintf("%d/%d",numerator/gcd,denominator/gcd)
	return
}


func getLowestCommonMultiple(a,b int) int{
	if a<b{
		a,b = b,a
	}
	i:=1
	for i<=b{
		if a*i%b == 0{
			return a*i
		}
		i++
	}
	return a*i
}

func getGreatestCommonDivisor(a, b int) int {
	if a < b {
		a ,b = b ,a
	}

	if a%b == 0 {
		return b
	}
	return getGreatestCommonDivisor(a%b,b)
}
