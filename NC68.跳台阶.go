package main

import "fmt"

func junp(num int) int {
	if num > 2 {
		return junp(num-1) + junp(num-2)
	} else if num > 1 {
		return 2
	}
	return 1
}
func main() {
	sum := 0
	fmt.Scan(&sum)
	fmt.Println(junp(sum))
}
