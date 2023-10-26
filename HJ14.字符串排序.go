package main

import (
	"fmt"
	"sort"
)

type Strs []string

func (T Strs) Len() int {
	return len(T)
}
func (T Strs) Less(i, j int) bool {
	return T[i] < T[j]
}
func (T Strs) Swap(i, j int) {
	T[i], T[j] = T[j], T[i]
}
func main() {
	nub := 0
	fmt.Scan(&nub)
	var strs Strs
	for i := 0; i < nub; i++ {
		str := ""
		fmt.Scan(&str)
		strs = append(strs, str)
	}
	sort.Sort(strs)
	for _, v := range strs {
		fmt.Println(v)
	}
}

//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	nub := 0
//	fmt.Scan(&nub)
//	strs := []string{}
//	for i := 0; i < nub; i++ {
//		str := ""
//		fmt.Scan(&str)
//		strs = append(strs, str)
//	}
//	sort.Strings(strs)
//	for _, v := range strs {
//		fmt.Println(v)
//	}
//}
