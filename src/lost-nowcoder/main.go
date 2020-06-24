package main

import (
	"fmt"
	"strings"
)

func main() {
	dire := [...]string{"N","W","S","E","N"}
	turn := 0
	fmt.Scan(&turn)
	var turnTmp string
	fmt.Scan(&turnTmp)
	turnTmp2 := strings.Split(turnTmp,"")
	tmp:=0
	for i := 0; i < turn; i++ {
		if turnTmp2[i] == "L" {
			tmp++
		}else{
			tmp--
		}
	}
	tmp = (tmp%4 +4)%4;
	fmt.Println(dire[tmp],tmp)
}
