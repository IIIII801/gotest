package main

import (
	"fmt"
)

func main() {
	a := 0
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a)
		}
	}
}
