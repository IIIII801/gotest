package main

import "fmt"

func main() {
	str := ""
	fmt.Scan(&str)
	m := make(map[int32]int32, len(str))
	for _, v := range str {
		if _, ok := m[v]; !ok && 0 <= v && v < 128 {
			m[v] = 0
		}
	}
	fmt.Println(len(m))
}
