package main

import "fmt"

func main() {
	s1 := "leetcode"
	s2 := "abc"
	fmt.Printf("Input：%s,result：%v \n",s1,isUnique(s1))
	fmt.Printf("Input：%s,result：%v \n",s2,isUnique(s2))
}

func isUnique(astr string) bool {
	str := []rune(astr)
	strMap := make(map[rune]int)
	for i , ch := range str{
		_,ok := strMap[ch]
		if ok {
			return false
		}
		strMap[ch] = i
	}
	return true
}