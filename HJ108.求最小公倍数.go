package main

import (
	"fmt"
	"math"
)

func main() {
	i, j := 0, 0
	fmt.Scan(&i, &j)
	tmp := geta(i, j)
	fmt.Println((i) * (j / tmp))
}
func geta(i, j int) int {
	for {
		if i-j == 0 {
			return i
		} else {
			i, j = int(math.Max(float64(i), float64(j))), int(math.Min(float64(i), float64(j)))
			i, j = i-j, j
		}
	}
}
