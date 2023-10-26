package main

import "fmt"

func main() {
	ans := [500]int{}
	num := 0
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		a := 0
		fmt.Scan(&a)
		ans[a]++
	}
	for i, v := range ans {
		if v > 0 {
			fmt.Println(i)
		}
	}

}
