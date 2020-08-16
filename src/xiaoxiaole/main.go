package main

import (
	"fmt"
	"strings"
)

func main() {
	var n,m int
	fmt.Scan(&n,&m)
	strs := make([]string,n)
	for i:=0;i<n;i++{
		fmt.Scan(&strs[i])
	}
	for _,v:= range strs{
		var res []string
		str2 := strings.Split(v,"")
		for i:=0;i<m;i++{
			res = checkstr(str2)
		}
		strRes := strings.Join(res,"")
		fmt.Println(strRes)
	}

}

func checkstr(str []string) []string {
	var x []int
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			x = append(x, i)
		}
	}
	var res []string = str
	i:=0
	for _, v := range x {
		res = append(res[:v-i],res[v+2-i:]...)
		i = i+2
	}
	return res
}
