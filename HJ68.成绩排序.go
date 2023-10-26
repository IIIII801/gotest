package main

import (
	"fmt"
	"sort"
)

type pepo struct {
	name string
	s    int
	flag int
}

func main() {
	var num, sor int
	fmt.Scan(&num, &sor)
	mylist := []pepo{}
	for i := 0; i < num; i++ {
		var name string
		var s int
		var tmp pepo
		fmt.Scan(&name, &s)
		tmp.name = name
		tmp.s = s
		tmp.flag = i
		mylist = append(mylist, tmp)
	}
	if sor == 1 {
		sort.Slice(mylist, func(i, j int) bool {
			if mylist[i].s != mylist[j].s {
				return mylist[i].s < mylist[j].s
			} else {
				return mylist[i].flag < mylist[j].flag
			}
		})
	} else {
		sort.Slice(mylist, func(i, j int) bool {
			if mylist[i].s != mylist[j].s {
				return mylist[i].s > mylist[j].s
			} else {
				return mylist[i].flag < mylist[j].flag
			}
		})
	}
	for _, v := range mylist {
		fmt.Println(v.name, v.s, v.flag)
	}
}
