package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("0","0"))
	fmt.Println(addBinary("11","1"))
	fmt.Println(addBinary("1010","1011"))
}

func addBinary(a, b string) string {
	sum := 0
	flag:= 0
	lenA := len(a)
	lenB := len(b)
	if lenA < lenB{
		lenA,lenB = lenB,lenA
		a,b = b,a
	}
	result := make([]byte , lenA + 1)

	for lenB > 0 {
		lenA --
		lenB --
		sum = int(a[lenA] - '0') + int(b[lenB] - '0') + flag
		flag = sum / 2
		sum = sum % 2
		result[lenA + 1] = byte(sum + '0')
	}
	for lenA > 0 {
		lenA --
		sum = int(a[lenA] - '0') + flag
		flag = sum / 2
		sum = sum % 2
		result[lenA + 1] = byte(sum + '0')
	}
	result[0] = byte(flag + '0')

	for lenA < len(result)-1 {
		if result[lenA] == '0' {
			lenA++
		} else {
			break
		}
	}
	return string(result[lenA:])
}
