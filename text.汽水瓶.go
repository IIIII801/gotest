package main

import "fmt"

func main() {
	for {
		num1 := 0
		num2 := 0
		fmt.Scan(&num1)
		if num1 == 0 {
			break
		}
		if num1 != 0 {
			for {
				if num1 < 2 {
					fmt.Println(num2)
					break
				} else if num1 > 2 {
					num2 += num1 / 3
					num1 = num1/3 + num1%3
				} else {
					num2 += 1
					num1 = 0
				}
			}
		}
	}
}
