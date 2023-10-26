package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ""
	fmt.Scan(&str)
	cut := [26]int{}
	for _, v := range str {
		cut[v-97]++
	}
	del := []int{}
	min := 20
	for i, v := range cut {
		if (v != 0) && (v <= min) {
			if v < min {
				min = v
				del = []int{i}
			} else {
				del = append(del, i)
			}
		}
	}
	for _, v := range del {
		str = strings.ReplaceAll(str, string(v), "")
		//strs := strings.Split(str, string(v+97))
		//str = strings.Join(strs, "")
	}
	fmt.Println(str)
}
