package main

import "fmt"

func main() {
	fmt.Println(minWindow("1034516327", "123"))
}
func minWindow(S string, T string) string {
	//write code here
	if len(S) == len(T) && S == T {
		return S
	}
	if len(S) < len(T) {
		return ""
	}
	var sign = make([]rune, 128)
	var tmp = make([]rune, 128)
	for _, val := range T {
		sign[val]++
	}
	var check func(arr []rune) bool
	check = func(arr []rune) bool {
		for i := 0; i < 128; i++ {
			if arr[i] < sign[i] {
				return false
			}
		}
		return true
	}
	l, r := 0, 0
	result := ""
	min := len(S)
	for r < len(S) {
		if sign[S[r]] != 0 {
			tmp[S[r]]++
		}
		for check(tmp) {
			if min > r-l+1 {
				min = r - l + 1
				result = S[l : r+1]
			}
			if sign[S[l]] != 0 {
				tmp[S[l]]--
			}
			l++
		}
		r++
	}
	return result

	//minr := 0
	//minl := 0
	//min := len(S)
	//ischange := false
	//hash := make(map[rune]int, len(T))
	//for _, v := range T {
	//	hash[v]++
	//}
	//tmp := make(map[rune]int, len(T))
	//for i := 0; i < len(S); i++ {
	//	if _, ok := hash[rune(S[i])]; ok {
	//		star, end := 0, len(S)-1
	//		var flag int
	//		tmp = cl
	//		for j := i; j < len(S); j++ {
	//			flag = -1
	//			if _, ok1 := hash[rune(S[j])]; ok1 {
	//				tmp[rune(S[j])]++
	//				for k, _ := range hash {
	//					if tmp[k] < hash[k] {
	//						flag = 0
	//						break
	//					} else {
	//						flag = 1
	//					}
	//				}
	//				if flag == 1 {
	//					star = i
	//					end = j
	//					ischange = true
	//					break
	//				}
	//
	//			}
	//		}
	//		if min-1 >= end-star && ischange {
	//			minl, minr = star, end
	//			min = end - star
	//		}
	//	}
	//}
	//if min != minr-minl {
	//	return ""
	//} else {
	//	return S[minl : minr+1]
	//}

}
