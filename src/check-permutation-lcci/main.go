package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
	fmt.Println(CheckPermutation("abc", "bad"))
}

func CheckPermutation(s1 string, s2 string) bool {
	str1 := strings.Split(s1,"")
	str2 := strings.Split(s2,"")
	sort.Strings(str1)
	sort.Strings(str2)
	tmp1 := strings.Join(str1,"")
	tmp2 := strings.Join(str2,"")
	return tmp2 == tmp1
}