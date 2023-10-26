package main

import (
	"fmt"
	"math/big"
)

func main() {
	aa := 5678
	bb := big.NewInt(int64(aa))
	aaa := bb.ProbablyPrime(0)
	fmt.Println(aaa)
}
