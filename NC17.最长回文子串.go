package main

import (
	"fmt"
)

func main() {
	s := ""
	fmt.Scan(&s)
	fmt.Println(getLongestPalindrome(s))
}
func getLongestPalindrome(A string) int {
	// write code here
	max := 1
	var ans func(i1, i2, d int)
	ans = func(i, i2, d int) {

		if i-d <= 0 || i2+d >= len(A)-1 || A[i-1-d] != A[d+i2+1] {
			if max < 1+2*d+i2-i {
				max = 1 + 2*d + i2 - i
			}
		} else {
			ans(i, i2, d+1)
		}
	}
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			ans(i, i+1, 0)
			ans(i, i, 0)
		} else {
			ans(i, i, 0)
		}
	}
	return max
}
