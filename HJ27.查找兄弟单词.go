package main

import (
	"fmt"
	"sort"
)

func main() {
	nub := 0
	fmt.Scan(&nub)
	brostr := []string{}
	for i := 0; i < nub; i++ {
		tmpstr := ""
		fmt.Scan(&tmpstr)
		brostr = append(brostr, tmpstr)
	}
	xstr := ""
	index := 0
	fmt.Scan(&xstr, &index)
	index--
	xstrsort := xstr
	tmp := []byte(xstrsort)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	xstrsort = string(tmp)
	laststr := []string{}
	for _, v := range brostr {
		v2 := v
		tmp := []byte(v2)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		v2 = string(tmp)
		if (v != xstr) && (v2 == xstrsort) {
			laststr = append(laststr, v)
		}
	}
	sort.Strings(laststr)
	fmt.Println(len(laststr))
	if len(laststr) > index {
		fmt.Println(laststr[index])
	}
}

//
//package main
//
//import (
//"fmt"
//"sort"
//)
//
//type str []byte
//
//func (t str) Len() int {
//	return len(t)
//}
//func (t str) Less(i, j int) bool {
//	return t[i] < t[j]
//}
//func (t str) Swap(i, j int) {
//	t[i], t[j] = t[j], t[i]
//}
//func main() {
//	nub := 0
//	fmt.Scan(&nub)
//	brostr := []string{}
//	for i := 0; i < nub; i++ {
//		tmpstr := ""
//		fmt.Scan(&tmpstr)
//		brostr = append(brostr, tmpstr)
//	}
//	xstr := ""
//	index := 0
//	fmt.Scan(&xstr, &index)
//	index--
//	xstrsort := xstr
//	tmp := str(xstrsort)
//	sort.Sort(tmp)
//	xstrsort = string(tmp)
//	laststr := []string{}
//	for _, v := range brostr {
//		v2 := v
//		tmp := str(v2)
//		sort.Sort(tmp)
//		v2 = string(tmp)
//		if (v != xstr) && (v2 == xstrsort) {
//			laststr = append(laststr, v)
//		}
//	}
//	sort.Strings(laststr)
//	fmt.Println(len(laststr))
//	if len(laststr) > index {
//		fmt.Println(laststr[index])
//	}
//}
