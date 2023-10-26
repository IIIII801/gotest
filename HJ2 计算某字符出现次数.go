package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Bytes()
	scan.Scan()
	byte1 := scan.Text()
	ans := 0
	if byte1[0] >= 'a' && byte1[0] <= 'z' {
		for _, v := range str {
			if v == byte1[0] || v+('a'-'A') == byte1[0] {
				ans++
			}
		}
	} else if byte1[0] >= 'A' && byte1[0] <= 'Z' {
		for _, v := range str {
			if v == byte1[0] || v-('a'-'A') == byte1[0] {
				ans++
			}
		}
	} else {
		for _, v := range str {
			if v == byte1[0] {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
