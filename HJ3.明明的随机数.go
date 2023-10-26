package main

import (
	"fmt"
	"sort"
)

func main() {
	numsum := 0
	fmt.Scan(&numsum)
	m := make(map[int]int, numsum)
	for i := 0; i < numsum; i++ {
		num := 0
		fmt.Scan(&num)
		if _, ok := m[num]; !ok {
			m[num] = i
		}
	}
	nus := make([]int, 0)
	for v := range m {
		nus = append(nus, v)
	}
	sort.Ints(nus)
	for _, v := range nus {
		fmt.Println(v)
	}
}
