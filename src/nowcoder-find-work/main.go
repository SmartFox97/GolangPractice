package main

import (
	"fmt"
	"sort"
)

type work struct {
	di int
	pi int
}

type workList []work

var works = []work{}

func setWorks(n int) {
	for i:=0;i<n;i++{
		di,pi := 0,0
		fmt.Scan(&di ,&pi)
		tmp := work{
			di: di,
			pi: pi,
		}
		works = append(works,tmp)
	}
}


func main() {
	var n,m int
	fmt.Scan(&n,&m)
	setWorks(n)
	var friends = make([]int,m)
	for i := 0; i < m; i++ {
		fmt.Scan(&friends[i])
	}
	sort.Sort(workList(works))
	for _,v := range friends{
		tmp := 0
		for _,pi := range works{
			if pi.di <= v{
				if tmp <= pi.pi{
					tmp = pi.pi
				}
			}else {
				break
			}
		}
		fmt.Println(tmp)
	}
}

func (myWork workList) Len() int {
	return len(myWork)
}
func (myWork workList) Less(i, j int) bool {
	return myWork[i].di < myWork[j].di
}
func (myWork workList) Swap(i, j int) {
	myWork[i],myWork[j] = myWork[j],myWork[i]
}