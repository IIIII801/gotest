package main

import (
	"fmt"
	"math/big"
)

func main() {
	ss := 0
	fmt.Scan(&ss)
	for i := ss / 2; i > 0; i++ {
		tmp1 := big.NewInt(int64(i))
		tmp2 := big.NewInt(int64(ss - i))
		if tmp1.ProbablyPrime(0) && tmp2.ProbablyPrime(0) {
			fmt.Println(ss - i)
			fmt.Println(i)
			break
		}
	}
}
